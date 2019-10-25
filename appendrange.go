package piscine

func AppendRange(s, b int) []int {
	var answer []int

	for i := s; i < b; i++ {

		answer = append(answer, i)
	}
	return answer
}
