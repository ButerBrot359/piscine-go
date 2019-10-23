package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {

	if n == 0 {
		z01.PrintRune('0')
	} else {
		s := []int{}
		s = CommonBoy(n, s)
		s = SortInt(s)
		for _, v := range s {
			z01.PrintRune(rune(v) + 48)
		}
	}

}
func CommonBoy(n int, s []int) []int {
	if n == 0 {
		return s
	} else {
		s = append(s, n%10)
		s = CommonBoy(n/10, s)
	}
	return s
}
func SortInt(s []int) []int {
	nb := StrLen(s)
	for i := 0; i < nb; i++ {
		for j := i; j < nb; j++ {
			if s[i] > s[j] {
				Swap(s, i, j)
			}
		}
	}
	return s
}
func StrLen(s []int) int {
	m := 0
	for range s {
		m++
	}
	return m
}
func Swap(s []int, i, j int) {
	temp := s[i]
	s[i] = s[j]
	s[j] = temp
}
