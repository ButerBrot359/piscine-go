package piscine

func BasicJoin(s []string) string {
	b := ""
	for _, v := range s {
		b += v
	}
	return b
}
