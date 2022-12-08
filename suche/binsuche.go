package suche

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
