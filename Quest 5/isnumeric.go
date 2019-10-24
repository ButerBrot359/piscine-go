package piscine

func IsNumeric(str string) bool {
	m := 0
	for range str {
		m++
	}
	nbr := []rune(str)
	for i := 0; i < m; i++ {
		if nbr[i] >= '0' && nbr[i] <= '9' {
			continue
		} else {
			return false
		}
	}
	return true
}
