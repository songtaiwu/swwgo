package cert

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"io/ioutil"
)

// PemBytesToFile 把pem编码的证书 写入到文件
func PemBytesToFile(filePath string, pemBytes []byte) error {
	err := ioutil.WriteFile(filePath, pemBytes, 0644)
	if err != nil {
		return err
	}
	return nil
}

// FileToPemBytes 把文件中的内容，转为pem证书
func FileToPemBytes(filePath string) ([]byte, error) {
	certBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return certBytes, nil
}

// --------------------------------------------------
// ----------------- 从pem文件中读取  -----------------
// --------------------------------------------------

// GetCsrFromPemFile
// 从csr的pem编码的文件中读取，转为*x509.CertificateRequest
func GetCsrFromPemFile(path string) (*x509.CertificateRequest, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(file)
	if block == nil {
		return nil, errors.New("证书错误")
	}

	request, err := x509.ParseCertificateRequest(block.Bytes)
	if err != nil {
		return nil, err
	}
	return request, nil
}


// GetCertFromPemFile
// 从pem证书文件读取，转为*x509.Certificate
func GetCertFromPemFile(path string) (*x509.Certificate, error) {
	certBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(certBytes)
	if block == nil {
		return nil, errors.New("证书错误")
	}

	x509Cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, err
	}

	return x509Cert, err
}

// GetPublicKeyFromPemFile
// 从pem证书文件读取出公钥
func GetPublicKeyFromPemFile(path string) (*rsa.PublicKey, error) {
	certBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(certBytes)
	if block == nil {
		return nil, errors.New("证书错误")
	}

	x509Cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, err
	}

	publicKey, ok := x509Cert.PublicKey.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("证书类型非RAS的公钥")
	}
	return publicKey, err
}

// GetPrivateKeyFromPemFile
// 从pem私钥文件读取出私钥
func GetPrivateKeyFromPemFile(path string) (*rsa.PrivateKey, error) {
	certBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(certBytes)
	if block == nil {
		return nil, errors.New("证书错误")
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return privateKey, err
}