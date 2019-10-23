package piscine

func NRune(s string, n int) rune {
	b := []rune(s)
	return b[n-1]
}
