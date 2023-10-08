package core

func QuantizeBlock(binaryData string) string {
	leadingZeros := ""
	for i := 0; i < 64-len(binaryData); i++ {
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
