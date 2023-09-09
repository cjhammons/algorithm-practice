package bubblesort

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	tests := []struct {
		input  []int
		output []int
	}{
		{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{[]int{2, 7, 9, 5, 3, 3, 8, 9, 6, 5, 3, 2, 0, 6}, []int{0, 2, 2, 3, 3, 3, 5, 5, 6, 6, 7, 8, 9, 9}},
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
	}

	for i, test := range tests {
		bubbleSort(test.input)
		if !reflect.DeepEqual(test.input, test.output) {
			t.Errorf("Test case %d failed. Got: %v, Expected: %v", i+1, test.input, test.output)
		}
	}
}
