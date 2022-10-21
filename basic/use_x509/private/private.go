package private

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
)

// 生成加密的私钥, 输出为pem编码的数据
// 1. 生成私钥
// 2. 将其转换为PEM格式
// 3. 加密PEM
func GenPriToEncryptedPEM(bits int, pwd string) ([]byte, error) {
	// 按照bits指定的rsa长度生成私钥
	key, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, err
	}

	// 转为pem格式
	// Bytes属性对应的是DER编码内容
	block := &pem.Block{
		Type: "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	}

	// 对pem加密
	if pwd != "" {
		block, err = x509.EncryptPEMBlock(rand.Reader, block.Type, block.Bytes, []byte(pwd), x509.PEMCipherAES256)
		if err != nil {
			return nil, err
		}
	}

	return pem.EncodeToMemory(block), err
}

// 生成RSA私钥，输出到文件
func GenRSAPriToFile(fileName, passwd string, bits int) error {
	key, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return err
	}

	data := x509.MarshalPKCS1PrivateKey(key)
	err = encodePrivPemFile(fileName, passwd, data)
	return err
}

// 把私钥输出到文件中
// fileName 要输出的文件名称
// password 用于给私钥加密的密码
// data 私钥的DER编码格式
func encodePrivPemFile(fileName, password string, data []byte) error {
	block, err := x509.EncryptPEMBlock(rand.Reader, "RSA PRIVATE KEY", data, []byte(password), x509.PEMCipherAES256)
	if err != nil {
		return err
	}

	file, err := os.Create(fileName)
	if err != nil {
		return err
	}

	err = pem.Encode(file, block)
	if err != nil {
		return err
	}
	return nil
}