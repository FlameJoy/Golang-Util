package util

import (
	"fmt"
	"math/rand"
	"time"
)

// Generate random filename
func FilenameGen(length int) string {
	source := rand.NewSource(time.Now().Unix())
	rnd := rand.New(source)
	randBytes := make([]byte, length)
	rnd.Read(randBytes)
	filename := fmt.Sprintf("%x%d", randBytes, time.Now().Unix())
	return filename
}
