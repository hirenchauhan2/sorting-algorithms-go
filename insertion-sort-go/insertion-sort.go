package insertionsort

// InsertionSort function implements the insertion sort algorithm
func InsertionSort(list []int) {
	// get the count of elements in the list
	n := len(list)

	for i := 1; i < n; i++ {
		// select a value to be inserted
		valueToInsert := list[i]
		holePos := i

		// locate hole position for the element to be inserted
		for ; holePos > 0 && list[holePos-1] > valueToInsert; holePos-- {
			swap(list, holePos, holePos-1)
		}
	}
}

func swap(list []int, i, j int) {
	temp := list[i]
	list[i] = list[j]
	list[j] = temp
}
