package selectionsort

import "fmt"

// SelectionSort function implements the Selection sort algorithm
func SelectionSort(list []int) {
	// get the length of slice
	n := len(list)

	for i := 0; i < n-1; i++ {
		// set the current element as min
		min := i

		// check for minimum element in the list
		for j := i + 1; j < n; j++ {
			if list[j] < list[min] {
				min = j
			}
		}

		// if we get the new min, swap it with current element
		if min != i {
			fmt.Println("\nIteration:", i,
				"\nMin is", list[min], " at index:", min,
				"\nBefore Swapping: ", list)
			swap(list, min, i)
			fmt.Println("Iteration:", i, "\nAfter Swapping: ", list)

		}
	}
}

func swap(list []int, i, j int) {
	temp := list[i]
	list[i] = list[j]
	list[j] = temp
}
