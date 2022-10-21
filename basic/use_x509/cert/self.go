package cert

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"errors"
	"io/ioutil"
	"math/big"
	"time"
)

// GenSelfSignedCA 创建自签CA证书
// 返回
//	1、证书PEM编码，
//  2、私钥PEM编码
//	3、error
func GenSelfSignedCA() ([]byte, []byte, error){
	// 创建私钥
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, err
	}

	template := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			CommonName: "sww Root CA",
		},
		SignatureAlgorithm: x509.SHA256WithRSA,
		PublicKey: x509.RSA,
		//证书的开始时间
		NotBefore: time.Now(),
		//证书的结束时间
		NotAfter: time.Now().Add(time.Hour * 24 * 365),
		//证书用途
		KeyUsage: x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		//基本的有效性约束
		BasicConstraintsValid: true,
		//是否是根证书
		IsCA: true,
	}

	// 创建证书
	// 自签ca，所以template和parent参数一样
	certDer, err := x509.CreateCertificate(rand.Reader, &template, &template, &privateKey.PublicKey, privateKey)
	if err != nil {
		return nil, nil, err
	}

	// 证书PEM编码
	certBuffer := bytes.Buffer{}
	certBlock := pem.Block{Type: "CERTIFICATE", Bytes: certDer}
	err = pem.Encode(&certBuffer, &certBlock)
	if err != nil {
		return nil, nil, err
	}

	// 私钥PEM编码
	keyBuffer := bytes.Buffer{}
	keyBlock := pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(privateKey)}
	err = pem.Encode(&keyBuffer, &keyBlock)
	if err != nil {
		return nil, nil, err
	}

	return certBuffer.Bytes(), keyBuffer.Bytes(), nil
}

// --------------------------------------------------
// ----------------- 公钥加密，私钥解密 -----------------
// --------------------------------------------------

// PublicPemEncrypt 公钥加密
// plainText 需要加密的信息
// path 公钥证书地址，证书是pem编码
func PublicPemEncrypt(plainText []byte, path string) ([]byte, error) {
	// 读取文件，pem格式的
	publicCert, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// 转为block对象
	block, _ := pem.Decode(publicCert)

	// cer格式证书 转为 *Certificate
	x509Cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, err
	}
	publicKey := x509Cert.PublicKey.(*rsa.PublicKey)

	// 对明文进行加密
	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, plainText)
	if err != nil {
		return nil, err
	}

	// 返回密文
	return cipherText, nil
}

// PrivatePemDecrypt 私钥加密
// cipherText 密文
// path 私钥证书地址，证书是pem编码
func PrivatePemDecrypt(cipherText []byte, path string) ([]byte, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// pem编码私钥转为 *pem.Block
	block, _ := pem.Decode(file)
	if block == nil {
		return nil, errors.New("私钥错误")
	}

	// Block中der编码证书 转为 *rsa.PrivateKey
	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	plainText, err := rsa.DecryptPKCS1v15(rand.Reader, key, cipherText)
	if err != nil {
		return nil, err
	}

	return plainText, nil
}

// --------------------------------------------------
// ----------------- 对终端证书签名    -----------------
// --------------------------------------------------
