package cert

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"time"
)

/**
创建自签ca证书， 核心就是调用x509.CreateCertificate()方法，template和parent参数一样，则代表是自签的。

创建一个证书csr，核心是调用 x509.CreateCertificateRequest(), 只要提前创建好私钥。
*/

// GenSignedClientPem
// 创建终端证书, 指定ca证书和私钥路径，证书直接被ca签名
// 返回
//	1、证书PEM编码，
//  2、私钥PEM编码
//	3、error
func GenSignedClientPem(cn string, caCertPath, caKeyPath string) ([]byte, []byte, error) {
	// 读取ca证书私钥
	caPrivateKey, err := GetPrivateKeyFromPemFile(caKeyPath)
	if err != nil {
		return nil, nil, err
	}
	// 读取ca证书
	caX509cert, err := GetCertFromPemFile(caCertPath)
	if err != nil {
		return nil, nil, err
	}

	// 创建私钥
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, err
	}

	template := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			CommonName: cn,
		},
		SubjectKeyId: []byte{1, 2, 3, 4, 6},
		//证书的开始时间
		NotBefore: time.Now(),
		//证书的结束时间
		NotAfter: time.Now().Add(time.Hour * 24 * 365),
		//证书用途
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:    x509.KeyUsageDigitalSignature,
		//基本的有效性约束
		BasicConstraintsValid: true,
		//是否是根证书
		IsCA: true,
	}

	// 创建证书
	// 自签ca，所以template和parent参数一样
	certDer, err := x509.CreateCertificate(rand.Reader, &template, caX509cert, &privateKey.PublicKey, caPrivateKey)
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

// GenClientCsr
// 创建csr，一次把私钥、证书CSR都生成了，可以把返回内容直接写到文件
// 返回
// 1、私钥，PEM编码
// 2、CSR，PEM编码
// 3、error
func GenClientCsr(cn string, keyBits int) ([]byte, []byte, error) {
	// 生成私钥
	privateKey, err := rsa.GenerateKey(rand.Reader, keyBits)
	if err != nil {
		return nil, nil, err
	}

	template := &x509.CertificateRequest{
		Subject: pkix.Name{
			//Organization: []string{"xxx"},
			//OrganizationalUnit: []string{"xxx"},
			CommonName: cn,
		},
		SignatureAlgorithm: x509.SHA256WithRSA,
		PublicKey:          x509.RSA,
	}
	// 生成csr，der编码格式
	request, err := x509.CreateCertificateRequest(rand.Reader, template, privateKey)
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

	// 证书请求PEM编码
	csrBuffer := bytes.Buffer{}
	csrBlock := pem.Block{Type: "CERTIFICATE REQUEST", Bytes: request}
	err = pem.Encode(&csrBuffer, &csrBlock)
	if err != nil {
		return nil, nil, err
	}

	return keyBuffer.Bytes(), csrBuffer.Bytes(), nil
}

// SignCsr
// 用ca证书签发一个csr
// 返回
// 1、证书，PEM编码
// 2、error
func SignCsr(csrPath string, caCertPath, caKeyPath string) ([]byte, error) {
	// 读取ca证书，转为*x509.Certificate
	caCert, err := GetCertFromPemFile(caCertPath)
	if err != nil {
		return nil, err
	}

	// 读取ca的私钥
	caPrivateKey, err := GetPrivateKeyFromPemFile(caKeyPath)
	if err != nil {
		return nil, err
	}

	// 读取csr文件，转为*x509.Certificate
	csr, err := GetCsrFromPemFile(csrPath)
	if err != nil {
		return nil, err
	}

	template := x509.Certificate{
		SerialNumber: big.NewInt(1),
		SubjectKeyId: []byte{1, 2, 3, 4, 6},
		//证书的开始时间
		NotBefore: time.Now(),
		//证书的结束时间
		NotAfter: time.Now().Add(time.Hour * 24 * 365),
		//证书用途
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:    x509.KeyUsageDigitalSignature,
		//基本的有效性约束
		BasicConstraintsValid: true,
		//是否是根证书
		IsCA: true,
	}
	certificate, err := x509.CreateCertificate(rand.Reader, &template, caCert, csr.PublicKey, caPrivateKey)
	if err != nil {
		return nil, err
	}

	// 证书PEM编码
	certBuffer := bytes.Buffer{}
	certBlock := pem.Block{Type: "CERTIFICATE", Bytes: certificate}
	err = pem.Encode(&certBuffer, &certBlock)
	if err != nil {
		return nil, err
	}

	return certBuffer.Bytes(), nil
}
