package linear_search

import (
	"testing"
)

func TestLinearSearch(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	target := 3
	if !linearSearch(arr, target) {
		t.Errorf("linearSearch(%v, %d) should be true", arr, target)
	}

	target = 6
	if linearSearch(arr, target) {
		t.Errorf("linearSearch(%v, %d) should be false", arr, target)
	}
}
