package util

import (
	"fmt"
	"math/rand"
	"time"
)

// Popular operators DEF-codes
var def = []int{901, 903, 904, 905, 906, 911, 921, 950, 961, 963, 965, 981}

// Telephone number generator
func TelGen(length int, mask bool) string {
	source := rand.NewSource(time.Now().Unix())
	rnd := rand.New(source)
	runes := []rune("1234567890")
	randRunes := make([]rune, length)
	for i := 0; i < len(randRunes); i++ {
		randRunes[i] = runes[rnd.Intn(len(runes))]
	}
	postfix := string(randRunes)
	defCode := def[rnd.Intn(len(def))]
	if mask {
		phone := fmt.Sprintf("+7(%d)-%s-%s-%s", defCode, postfix[:3], postfix[3:5], postfix[5:])
		return phone
	}
	phone := fmt.Sprintf("+7%d%s", defCode, postfix)
	return phone
}
