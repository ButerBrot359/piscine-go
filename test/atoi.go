package piscine

func Atoi(s string) int {

	res := 0
	for _, val := range s {
		if val < '0' {
			res = '-'
		}
		if val >= '0' && val <= '9' {
			a := 0
			for i := '1'; i <= val; i++ {
				a++
			}
			res = res*10 + a
		} else {
			res = 0
			return res
		}
	}
	return res
}	