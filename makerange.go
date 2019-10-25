package piscine

func MakeRange(min, max int) []int {
	t := max - min
	answer := make([]int, t)

	for i := min; i < max; i++ {
		answer[i] = i
	}
	return answer

}
