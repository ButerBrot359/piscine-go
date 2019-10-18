package piscine

func StrRev(b string) string{

	var cox string
	for _, geralt := range b {
		cox = string(geralt) + cox  
	}
	return cox
}
