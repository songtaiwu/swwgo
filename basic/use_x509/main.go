package main

import (
	"fmt"
	"log"
	"swwgo/basic/use_x509/cert"
)

func main() {
	// ------- 创建根CA证书
	//public, private, err := cert.GenSelfSignedCA()
	//if err != nil {
	//	fmt.Println(err)
	//}
	//cert.PemBytesToFile("ca.crt", public)
	//cert.PemBytesToFile("ca.key", private)


	// ------- 证书的加密和解密
	str := "hello world"
	encrypt, err := cert.PublicPemEncrypt([]byte(str), "ca.crt")
	if err != nil {
		log.Fatalln(err)
	}
	decrypt, err := cert.PrivatePemDecrypt(encrypt, "ca.key")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(decrypt))


	// ------- ca签客户端证书
	//pemPublic, pemBytes, err := cert.GenSignedClientPem("client123", "ca.crt", "ca.key")
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//cert.PemBytesToFile("client.crt", pemPublic)
	//cert.PemBytesToFile("client.key", pemBytes)
	
	// ------- 创建客户端证书csr 并签发
	//clientKey, clientCsr, err := cert.GenClientCsr("client999", 2048)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//cert.PemBytesToFile("client.key", clientKey)
	//cert.PemBytesToFile("client.csr", clientCsr)

	//clientCert, err := cert.SignCsr("client.csr", "ca.crt", "ca.key")
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//cert.PemBytesToFile("client.crt", clientCert)
}
