package main

import ("github.com/01-edu/z01"
	"os"
)

func main() {
	arg := os.Args
	for i, arg := range arg {
		if i != 0 {
			for _, c := range arg {
				z01.PrintRune(c)
			}
			z01.PrintRune('\n')
		}
		
	}
	
	
}