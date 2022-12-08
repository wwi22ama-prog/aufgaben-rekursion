package elemententfernen

// Liefert eine Liste, die alle Elemente aus list enthält,
// außer dem an Stelle pos.
func RemoveElement(list []int, pos int) []int {
	if len(list) == 0 || pos < 0 {
		return []int{}
	}

	head := list[:1]
	tail := list[1:]
	if pos == 0 {
		return tail
	}
	return append(head, RemoveElement(tail, pos-1)...)
}
