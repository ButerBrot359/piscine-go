package piscine

func StrLen(str string) int {
	b := 0
	for _, cc := range str {
		if cc == cc {
			b++
		}
	}
	return b
}
