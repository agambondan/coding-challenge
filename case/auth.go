package main

import (
	"crypto/rand"
	"encoding/base64"
	"log"
)

func main() {

	//password := "me1JPZ+bNQEEpSoItRkTqwZQSizjyf4efccLrw=="
	//
	//// decrypt password
	//decodeString, _ := base64.StdEncoding.DecodeString(password)
	//decrypt, err := lib.CipherDecrypt(decodeString, "AES_CIPHER_SECRET_KEY_KUDU_32BIT")
	//if err != nil {
	//	panic(err)
	//}
	//decryptSalt, err := lib.CipherDecrypt(decrypt, "PROTECT_YOUR_PASSWORD")
	//if err != nil {
	//	panic(err)
	//}
	//log.Println(string(decryptSalt))
	log.Println(GenerateKey())
}

func GenerateKey() string {
	ret := make([]byte, 32)

	if _, err := rand.Read(ret); err != nil {
		panic(err)
	}

	return base64.StdEncoding.EncodeToString(ret)
}
