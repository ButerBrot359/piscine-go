package piscine

func IterativePower(nb int, power int) int {
	res := nb
	if power <= 20 && power > 0 {
		for i := 1; i <= power-1; i++ {
			res *= nb
		}
	} else if power == 0 {
		return 1
	} else if power > 20 || power < 0 {
		return 0
	}
	return res
}
