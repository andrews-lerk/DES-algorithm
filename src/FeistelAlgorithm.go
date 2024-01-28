package src

import (
	"DES-algorithm/src/tables"
	"strconv"
)

func FeistelAlgorithm(leftBlock string, rightBlock string, lapKeys [16]string) string {
	for _, key := range lapKeys {
		extendedXoredBlock := xor([]rune(rightBlockPermutation([]rune(rightBlock))), []rune(key))
		resultRightBlock := xor([]rune(leftBlock), []rune(finalLapPermutation([]rune(blockReplacement(extendedXoredBlock)))))
		leftBlock = rightBlock
		rightBlock = resultRightBlock
	}
	return finalFeistelAlgorithmPermutation([]rune(rightBlock + leftBlock))
}

func blockReplacement(block string) string {
	result := ""
	for i := 0; i < len(block); i += 6 {
		SblockNumber := i / 6
		lineNumber, _ := strconv.ParseInt(block[i : i+6][:1]+block[i : i+6][5:], 2, 64)
		columnNumber, _ := strconv.ParseInt(block[i : i+6][1:5], 2, 64)
		targetValue := strconv.FormatInt(int64(tables.ReplacementMatrix[SblockNumber][lineNumber][columnNumber]), 2)
		result += QuantizeBlock(targetValue, 4)
	}
	return result
}

func rightBlockPermutation(bitsChain []rune) string {
	result48BitBlock := ""
	for _, v := range tables.ExtensionsRightBlockTable {
		result48BitBlock += string(bitsChain[v-1])
	}
	return result48BitBlock
}

func finalLapPermutation(bitsChain []rune) string {
	result48BitBlock := ""
	for _, v := range tables.FinalLapBlockPermutationTable {
		result48BitBlock += string(bitsChain[v-1])
	}
	return result48BitBlock
}

func finalFeistelAlgorithmPermutation(bitsChain []rune) string {
	result48BitBlock := ""
	for _, v := range tables.FinalFeistelAlgorithmPermutationTable {
		result48BitBlock += string(bitsChain[v-1])
	}
	return result48BitBlock
}

func xor(block1 []rune, block2 []rune) string {
	resultBlock := ""
	for i := 0; i < len(block2); i++ {
		if block1[i] != block2[i] {
			resultBlock += "1"
		} else {
			resultBlock += "0"
		}
	}
	return resultBlock
}
