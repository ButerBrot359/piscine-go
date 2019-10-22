package piscine

func IsPrime(nb int) bool {

	if nb == 2 {
		return true
	} else if nb > 1 {
		for i := 2; i < nb; i++ {
			if nb%i != 0 {
				return true
			} else if nb == i*i {
				return false
			} else {
				return false
			}
		}
	} else {
		return false
	}
	return false
}
