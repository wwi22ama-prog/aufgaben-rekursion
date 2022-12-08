package suche

func Search(list []int, key int) int {
	if len(list) == 0 {
		return 0
	}
	head := list[0]
	tail := list[1:]

	if head == key {
		return 0
	}
	return Search(tail, key) + 1
}
