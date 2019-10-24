package piscine

func IsAlpha(str string) bool {
	nb := Friday(str)
	modul := []rune(str)
	for i := 0; i < nb; i++ {
		if (modul[i] >= 'a' && modul[i] <= 'z') || (modul[i] >= 'A' && modul[i] <= 'Z') || (modul[i] >= '0' && modul[i] <= '9') {
			continue
		} else {
			return false
		}

	}
	return true
}
func Friday(s string) int {
	m := 0
	for range s {
		m++
	}
	return m
}
