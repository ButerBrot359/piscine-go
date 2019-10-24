package piscine

func Atoi(s string) int {
	nb := H(s)
	massiv := []rune(s)
	num := 0
	s = ""
	minus := 0
	plus := 0
	summa := 0
	for i := 0; i < nb; i++ {
		if massiv[i] >= '0' && massiv[i] <= '9' {
			num = 0
			for j := '0'; j < massiv[i]; j++ {
				num++
			}
			summa = summa*10 + num
		} else if massiv[i] == '-' && i == 0 {
			minus++
		} else if massiv[i] == '+' && i == 0 {
			plus++
		} else {
			return 0
		}
	}
	if minus == 1 {
		summa = -summa
	}
	return summa
}
func H(s string) int {
	m := 0
	for range s {
			m++
	}
	return m
}
