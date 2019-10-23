package piscine

func NRune(s string, n int) rune {
	b := []rune(s)
	nb := Len(s)
	if n > nb || n < 0 {
		return 0
	}
	return b[n-1]
}
func Len(s string) int {
	m := 0
	for range s {
		m++
	}
	return m
}
