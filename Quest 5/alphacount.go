package piscine

func AlphaCount(str string) int {
	count := 0
	for _, value := range str {
		if (value >= 'a' && value <= 'z') || (value >= 'A' && value <= 'Z') {
			count++
		} else {
			continue
		}

	}
	return count
}
