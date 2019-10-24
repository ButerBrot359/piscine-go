package piscine

func Capitalize(s string) string {

	nb := N(s)
	new := []rune(s)
	if new[0] >= 'a' && new[0] <= 'z' {
		new[0] = new[0] - 32
	}
	for i := 1; i < nb; i++ {
		if new[i] >= 'a' && new[i] <= 'z' && (new[i-1] < '0' || new[i-1] > '9' && new[i-1] < 'A' || new[i-1] > 'Z' && new[i-1] < 'a' || new[i-1] > 'z') {
			new[i] = new[i] - 32
		}
		if new[i] >= 'A' && new[i] <= 'Z' && (new[i-1] >= '0' && new[i-1] <= '9' || new[i-1] >= 'A' && new[i-1] <= 'Z' || new[i-1] >= 'a' && new[i-1] <= 'z') {
			new[i] = new[i] + 32
		}
	}
	return string(new)

}
func N(s string) int {
	m := 0
	for range s {
		m++
	}
	return m

}
