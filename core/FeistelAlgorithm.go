package core

import "fmt"

func FeistelAlgorithm(leftBlock string, rightBlock string, key string) {
	lapKeys := Create16LapKeys(key)
	fmt.Println(lapKeys)
}
