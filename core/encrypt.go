package core

import (
	"DES-algorithm/core/tables"
)

func Encrypt(text string, key string) {
	binaryData := Encode(text)
	binaryKey := QuantizeKey(Encode(key))
	for i := 0; i < len(binaryData); i += 64 {
		var block64bit string
		if len(binaryData[i:]) < 64 {
			block64bit = QuantizeBlock(binaryData[i:])
		} else {
			block64bit = binaryData[i : i+64]
		}
		initialPermutation(&block64bit)
		FeistelAlgorithm(block64bit[:32], block64bit[32:], binaryKey)
	}
}

func initialPermutation(block64bit *string) {
	originalBitsChain := []rune(*block64bit)
	*block64bit = ""
	for _, v := range tables.InitialPermutationTable {
		*block64bit += string(originalBitsChain[v-1])
	}
}
