package src

func QuantizeBlock(binaryData string, targetLen int) string {
	leadingZeros := ""
	for i := 0; i < targetLen-len(binaryData); i++ {
		leadingZeros += "0"
	}
	return leadingZeros + binaryData
}

func QuantizeKey(binaryKey string) string {
	if len(binaryKey) > 64 {
		panic("Key must be <= 8 chars")
	}
	leadingZeros := ""
	for i := 0; i < 64-len(binaryKey); i++ {
		leadingZeros += "0"
	}
	return leadingZeros + binaryKey
}

func ReverseLapKeysArray(lapKeys [16]string) [16]string {
	var result [16]string
	for i := 15; i >= 0; i -= 1 {
		result[i] = lapKeys[15-i]
	}
	return result
}
