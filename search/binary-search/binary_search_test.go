package binary_search

import (
	"testing"
)

func testBinarySearchHappy(t *testing.T) {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	target := 3
	if !binarySearch(arr, target) {
		t.Errorf("binarySearch(%v, %d) should be true", arr, target)
	}

}

func testBinarySearchSad(t *testing.T) {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	target := 420
	if binarySearch(arr, target) {
		t.Errorf("binarySearch(%v, %d) should be false", arr, target)
	}

}

func testBinarySearchRecursiveHappy(t *testing.T) {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	target := 3
	if !binarySearchRecursive(arr, target, 0, len(arr)) {
		t.Errorf("binarySearch(%v, %d) should be true", arr, target)
	}

}

func testBinarySearchRecursiveSad(t *testing.T) {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	target := 420
	if binarySearchRecursive(arr, target, 0, len(arr)) {
		t.Errorf("binarySearchRecursive(%v, %d) should be false", arr, target)
	}

}
