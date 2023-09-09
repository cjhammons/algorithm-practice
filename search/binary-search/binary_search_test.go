package binary_search

import (
	"testing"
)

func TestBinarySearchHappy(t *testing.T) {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	target := 5
	if !binarySearch(arr, target) {
		t.Errorf("binarySearch(%v, %d) should be true", arr, target)
	}

	target = 1
	if !binarySearch(arr, target) {
		t.Errorf("binarySearch(%v, %d) should be true", arr, target)
	}

	target = 9
	if !binarySearch(arr, target) {
		t.Errorf("binarySearch(%v, %d) should be true", arr, target)
	}
}

func TestBinarySearchSad(t *testing.T) {
	arr := []int{0, 1, 2, 3, 5, 6, 7, 8, 9}
	target := 4
	if binarySearch(arr, target) {
		t.Errorf("binarySearch(%v, %d) should be false", arr, target)
	}

}

func TestBinarySearchRecursiveHappy(t *testing.T) {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	target := 3
	if !binarySearchRecursive(arr, target, 0, len(arr)) {
		t.Errorf("binarySearch(%v, %d) should be true", arr, target)
	}

	target = 1
	if !binarySearchRecursive(arr, target, 0, len(arr)) {
		t.Errorf("binarySearch(%v, %d) should be true", arr, target)
	}

	target = 9
	if !binarySearchRecursive(arr, target, 0, len(arr)) {
		t.Errorf("binarySearch(%v, %d) should be true", arr, target)
	}

}

func TestBinarySearchRecursiveSad(t *testing.T) {
	arr := []int{0, 1, 2, 3, 5, 6, 7, 8, 9}
	target := 4
	if binarySearchRecursive(arr, target, 0, len(arr)) {
		t.Errorf("binarySearchRecursive(%v, %d) should be false", arr, target)
	}

}
