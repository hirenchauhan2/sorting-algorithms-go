package bubblesort

// SortNumbers function implents Bubble Sort Algorithm for sorting
// Currently, it takes the slice of integers
func SortNumbers(numbers []int) {
	// count elements in the array
	elCount := len(numbers)

	for {
		var swapped bool = false
		for i := 0; i < elCount-1; i++ {
			if numbers[i] > numbers[i+1] {
				swap(numbers, i, i+1)
				swapped = true
			}
		}

		if !swapped {
			break
		}
	}
}

func swap(numbers []int, i, j int) {
	temp := numbers[i]
	numbers[i] = numbers[j]
	numbers[j] = temp
}
