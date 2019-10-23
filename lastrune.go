package piscine

func LastRune(s string) rune {
	nb := LLL(s)
	b := []rune(s)
	return b[nb-1]
}
func LLL(s string) int {
	m := 0
	for range s {
		m++
	}
	return m
}
