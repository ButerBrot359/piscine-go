package piscine

import "fmt"

func PrintStr(str string) {

	stringi := []byte(str)

	for _, letter := range stringi {

		fmt.Printf("%c", letter)
	}
	fmt.Println()
}
