package elementeaddieren

// Liefert die Summe aller Elemente in list.
func AddElements(list []int) int {
	if len(list) == 0 {
		return 0
	}
	head := list[0]
	tail := list[1:]
	return head + AddElements(tail)
}
