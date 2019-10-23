package piscine

func TrimAtoi(s string) int {
	nb := StrLen(s)
	massiv := []rune(s)
	plus := 0
	minus := 0
	summa := 0
	for i := 0; i < nb; i++ {
		if massiv[i] >= '0' && massiv[i] <= '9' {
			num := 0
			for j := '0'; j < massiv[i]; j++ {
				num++
			}
			summa = summa*10 + num
		} else if massiv[i] == '-' && summa == 0 {
			minus++
		} else if massiv[i] == '+' && summa == 0 {
			plus++
		} else {
			continue
		}
	}
	if minus == 1 {
		summa = -summa
	}
	return summa

}
func StrLen(s string) int {

	m := 0
	for range s {
		m++
	}
	return m
}
