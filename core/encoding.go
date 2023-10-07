package core

import (
	"fmt"
	"strconv"
)

const maxSymbolSize = 16

func Encode(text string) string {
	binary := ""
	format := "%0" + strconv.Itoa(maxSymbolSize) + "b"
	for _, symbol := range text {
		binary += fmt.Sprintf(format, symbol)
	}
	return binary
}

func Decode(binary string) string {
	text := ""
	for i := 0; i < len(binary); i += maxSymbolSize {
		binarySymbol := binary[i : i+maxSymbolSize]
		symbolCode, _ := strconv.ParseInt(binarySymbol, 2, 64)
		text += string(symbolCode)
	}
	return text
}
