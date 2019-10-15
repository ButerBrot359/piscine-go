package isnegative

import "github.com/01-edu/z01"

func main() {

	var number1, number2, number3 int = 1, 0, -1
	if number1 >=0 {
		z01.PrintRune('F')
	} else if number1 < 0 {
			z01.PrintRune('T')
	}
	z01.PrintRune('\n')
	if number2 >=0 {
		z01.PrintRune('F')
	} else if number2 < 0 {
			z01.PrintRune('T')
	}
	z01.PrintRune('\n')
	if number3 >=0 {
		z01.PrintRune('F')
	} else if number3 < 0 {
			z01.PrintRune('T')
	}
	z01.PrintRune('\n')

}