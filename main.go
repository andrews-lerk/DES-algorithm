package main

import (
	"DES-algorithm/core"
)

func main() {
	data := "eternity"
	key := "alekos"
	core.Encrypt(data, key)
}
