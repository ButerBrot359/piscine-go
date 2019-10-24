package piscine

func IsPrintable(str string) bool {
	m := 0
	for range str {
		m++
	}
	nbr := []rune(str)
	for i := 0; i < m; i++ {
		if nbr[i] >= 32 && nbr[i] <= 126 {
			continue
		} else {
			return false
		}
	}
	return true
}
