package piscine

func SplitWhiteSpaces(str string) []string {
	var index int
	countWords := 1
	for c := range str {
		if isWhiteSpace(str[c]) {
			countWords++
		}
	}
	var word string
	answer := make([]string, countWords)
	for _, j := range str {
		if j != ' ' && j != '\n' && j != '\t' {
			word += string(j)
		} else if (j == ' ' || j == '\n' || j == '\t') && word != "" {
			answer[index] = word
			word = ""
			index++
		}
	}
	answer[index] = word //last word
	return answer
}

func isWhiteSpace(r byte) bool {
	if r == ' ' || r == '\n' || r == '\t' {
		return true
	}
	return false
}
