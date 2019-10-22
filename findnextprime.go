package piscine

func FindNextPrime(nb int) int {

	if nb <= 1 {
		return 2
	} else {
		for i := 1; i <= nb; i++ {
			if IsPrime(nb) == false {
				nb++
			} else if IsPrime(nb) == true {
				return nb
			}
		}
	}
	return nb
}
func IsPrime(nb int) bool {

	if nb <= 1 {
		return false
	} else {
		for i := 2; i < nb; i++ {
			if nb%i == 0 {
				return false
			}
		}
		return true
	}

}
