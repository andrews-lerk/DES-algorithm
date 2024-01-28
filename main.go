package main

import (
	"DES-algorithm/src"
	"fmt"
)

func main() {
	data := "secret information to encryption"
	key := "lerk2003"
	encryptedData := src.Encrypt(data, key)
	fmt.Println(encryptedData)
	decryptedData := src.Decrypt(encryptedData, key)
	fmt.Println(decryptedData)
}
