package piscine

func AppendRange(s, b int) []int {
	var answer []int

	for i := 0; i < (b - s); i++ {

		answer = append(answer, i+1)
	}
	return answer
}
