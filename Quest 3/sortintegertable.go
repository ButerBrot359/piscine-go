package piscine

func SortIntegerTable(table []int) {
	for i := 0; i < len(table); i++ {
		for j := i; j < len(table); j++ {
			if table[i] > table[j] {
				swap(table, i, j)
			}
		}
	}
}
func swap(table []int, i, j int) {
	temp := table[i]
	table[i] = table[j]
	table[j] = temp
}
