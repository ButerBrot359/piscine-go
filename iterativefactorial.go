package piscine

func IterativeFactorial(nb int) int {
	res := 1
	if nb >= 0 {
		for i := 1; i <= nb; i++ {
			res = res * i
		}
		return res
	} else {
		res = 0
		return res
	}
}
