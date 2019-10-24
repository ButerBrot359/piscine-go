package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	arg := os.Args
	nb := NB(arg)
	lol := []string(arg)
	for i := nb - 1; i > 0; i-- {
		for _, b := range lol[i] {
			z01.PrintRune(b)
		}
		z01.PrintRune('\n')
	}

}
func NB(s []string) int {
	m := 0
	for range s {
		m++
	}
	return m
}
