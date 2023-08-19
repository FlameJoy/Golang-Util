package util

import (
	"math/rand"
	"time"
)

// Generate password
func PswdGen(length int, symbols bool) string {
	source := rand.NewSource(time.Now().Unix())
	rnd := rand.New(source)
	runes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	runesWithSymbols := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!@#$%&*")
	pswdRand := make([]rune, length)
	for i := 0; i < length; i++ {
		if symbols {
			pswdRand[i] = runesWithSymbols[rnd.Intn(len(runesWithSymbols))]
		} else {
			pswdRand[i] = runes[rnd.Intn(len(runes))]
		}
	}
	pswd := string(pswdRand)
	return pswd
}
