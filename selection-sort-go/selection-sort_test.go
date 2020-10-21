package selectionsort

import (
	"fmt"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	list := []int{10, 25, 11, 3, 1, 2}
	listStr := fmt.Sprint(list)
	t.Log("Before sorting: ", listStr)
	SelectionSort(list)
	listStr2 := fmt.Sprint(list)
	t.Log("After sorting: ", listStr2)

	if listStr2 != "[1 2 3 10 11 25]" {
		t.Error("Sorting failed.",
			"\nOriginal: ", listStr,
			"\nExpected: [1 2 3 10 11 25]",
			"\nReceived: ", listStr2)
	}
}
