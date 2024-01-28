package src

import (
	"DES-algorithm/src/tables"
)

func Encrypt(text string, key string) string {
	binaryData := Encode(text)
	binaryKey := QuantizeKey(Encode(key))
	encryptedData := ""
	for i := 0; i < len(binaryData); i += 64 {
		var block64bit string
		if len(binaryData[i:]) < 64 {
			block64bit = QuantizeBlock(binaryData[i:], 64)
		} else {
			block64bit = binaryData[i : i+64]
		}
		lapKeys := Create16LapKeys(binaryKey)
		initialPermutation(&block64bit)
		encryptedBlock := FeistelAlgorithm(block64bit[:32], block64bit[32:], lapKeys)
		encryptedData += encryptedBlock
	}
	return encryptedData
}

func Decrypt(binaryData string, key string) string {
	binaryKey := QuantizeKey(Encode(key))
	decryptedData := ""
	for i := 0; i < len(binaryData); i += 64 {
		block64bit := binaryData[i : i+64]
		lapKeys := ReverseLapKeysArray(Create16LapKeys(binaryKey))
		initialPermutation(&block64bit)
		decryptedBlock := FeistelAlgorithm(block64bit[:32], block64bit[32:], lapKeys)
		decryptedData += decryptedBlock
	}
	return Decode(decryptedData)
}

func initialPermutation(block64bit *string) {
	originalBitsChain := []rune(*block64bit)
	*block64bit = ""
	for _, v := range tables.InitialPermutationTable {
		*block64bit += string(originalBitsChain[v-1])
	}
}
