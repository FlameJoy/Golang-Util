package main

import (
	"Golang-Util/util"
	"fmt"
)

func main() {
	fmt.Println(util.TelGen(7, true))
	// time.Sleep(time.Second * 1)
	fmt.Println(util.TelGen(7, false))
	fmt.Println(util.FilenameGen(12))
	fmt.Println(util.PswdGen(24, false))
	fmt.Println(util.PswdGen(24, true))
}
