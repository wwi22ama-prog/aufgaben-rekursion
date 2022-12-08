package listenvergleich

// Liefert true, falls die beiden Listen gleich sind.
func ListsEqual(list1, list2 []int) bool {
	if len(list1) != len(list2) {
		return false
	}
	if len(list1) == 0 {
		return len(list2) == 0
	}
	if len(list2) == 0 {
		return len(list1) == 0
	}

	head1 := list1[0]
	head2 := list2[0]
	tail1 := list1[1:]
	tail2 := list2[1:]

	if head1 != head2 {
		return false
	}
	return ListsEqual(tail1, tail2)
}
