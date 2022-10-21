package main

import "swwgo/basic/use_x509/private"

func main() {
	_ = private.GenRSAPriToFile("key.pem", "", 2048)
}
