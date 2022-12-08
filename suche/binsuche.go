package suche

// Liefert die Position von key in der Liste, falls key enthalten ist.
// Liefert die Länge der Liste, falls key nicht enthalten ist.
// Führt eine binäre Suche durch.
func BinSearch(list []int, key int) int {
	if len(list) == 0 {
		return 0
	}
	pos_mitte := len(list) / 2

	if key == list[pos_mitte] {
		return pos_mitte
	}
	if key < list[pos_mitte] {
		return BinSearch(list[:pos_mitte], key)
	}
	return BinSearch(list[pos_mitte+1:], key) + pos_mitte + 1
}
