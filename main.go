package main

import (
	"Golang-Util/util"
	"fmt"
)

func main() {
	fmt.Println(util.RandTelGen(7, true))
	fmt.Println(util.RandFilenameGen(12))
}
