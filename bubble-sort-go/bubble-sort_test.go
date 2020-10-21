package bubblesort

import (
	"fmt"
	"testing"
)

func TestSort(t *testing.T) {
	var numbers []int = []int{5, 4, 6, 3, 2, 1}
	numsStr := fmt.Sprint(numbers)
	t.Log("Before sort", numsStr)
	SortNumbers(numbers)
	numsStr2 := fmt.Sprint(numbers)
	t.Log("After sort", numsStr2)
	if numsStr2 != "[1 2 3 4 5]" {
		t.Error("The sorting failed. \n",
			"Original: ", numsStr,
			"\nExpecting: [1 2 3 4 5]",
			"\nReceived: ", numsStr2)
	}
}
