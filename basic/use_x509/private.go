package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
)

// 生成加密的私钥
// 1. 生成私钥
// 2. 将其转换为PEM格式
// 3. 加密PEM
func PrivateKeyToEncryptedPEM(bits int, pwd string) ([]byte, error) {
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
