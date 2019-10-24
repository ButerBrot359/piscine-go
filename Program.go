package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	arg := os.Args
	for _, i := range arg[1] {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')
}
