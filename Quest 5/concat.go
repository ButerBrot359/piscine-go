package piscine

func Concat(str1 string, str2 string) string {
	n := []string(str1)
	n2 := []string(str2)
	c := append(n, n2)
	return c
}
