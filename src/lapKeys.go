package src

import (
	"DES-algorithm/src/tables"
)

func Create16LapKeys(key string) [16]string {
	keyPermutation(&key)
	var lapKeys [16]string
	for i, v := range tables.CycleShiftTable {
		cycleShift(&key, v)
		lapKeys[i] = createLapKey(&key)
	}
	return lapKeys
}

func keyPermutation(key *string) {
	originalBitsChain := []rune(*key)
	*key = ""
	for _, v := range tables.KeyPermutationTable {
		*key += string(originalBitsChain[v-1])
	}
}

func cycleShift(key *string, shiftValue int) {
	leftPart := []rune(*key)[:28]
	rightPart := []rune(*key)[28:]
	leftPartShifted := append(leftPart[shiftValue:], leftPart[:shiftValue]...)
	rightPartShifted := append(rightPart[shiftValue:], rightPart[:shiftValue]...)
	shiftedKey := append(leftPartShifted, rightPartShifted...)
	*key = ""
	for _, v := range shiftedKey {
		*key += string(v)
	}
}

func createLapKey(key *string) string {
	originalBitsChain := []rune(*key)
	lapKey := ""
	for _, v := range tables.FinalLapKeyPermutationTable {
		lapKey += string(originalBitsChain[v-1])
	}
	return lapKey
}
