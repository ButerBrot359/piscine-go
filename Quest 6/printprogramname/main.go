package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	arg := os.Args
	for _, i := range arg[0] {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')
}
