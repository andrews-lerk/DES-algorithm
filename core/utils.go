package core

func QuantizeBlock(binaryData string) string {
	needLength := 64 - len(binaryData)
	for i := 0; i < needLength; i++ {
		binaryData += "0"
	}
	return binaryData
}
