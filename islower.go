package piscine

func IsLower(str string) bool {
	m := 0
	for range str {
		m++
	}
	nbr := []rune(str)
	for i := 0; i < m; i++ {
		if nbr[i] >= 'a' && nbr[i] <= 'z' {
			continue
		} else {
			return false
		}
	}
	return true
}
