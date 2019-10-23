package piscine

func Atoi(s string) int {
	nb := StrLen(s)
	massiv := []rune(s)
	num := 0
	s = ""
	minus := 0
	plus := 0
	summa := 0
	for i := 0; i < nb; i++ {
		num = 0
		if massiv[i] >= '0' && massiv[i] <= '9' {
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
