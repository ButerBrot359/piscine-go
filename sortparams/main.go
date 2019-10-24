package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	arg := os.Args
	lol := []string(arg)
	nb := Batyr(arg)
	for i := 0; i < nb; i++ {
		for j := i; j < nb; j++ {
			if lol[i] > lol[j] {
				SwapBabe(lol, i, j)
			}
		}
	}
	for i := range arg {
		if i != 0 {
			for _, c := range lol[i] {
				z01.PrintRune(c)
			}
			z01.PrintRune('\n')
		}
	}
}
func Batyr(s []string) int {
	m := 0
	for range s {
		m++
	}
	return m
}
func SwapBabe(b []string, i, j int) {
	temp := b[i]
	b[i] = b[j]
	b[j] = temp
}
