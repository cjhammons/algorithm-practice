package binary_search

func binarySearch(arr []int, v int) bool {
	lo := 0
	hi := len(arr)

	for lo < hi {
		m := int(lo + (hi-lo)/2)
		n := arr[m]

		if v == n {
			return true
		} else if n > v {
			hi = m
		} else {
			lo = m + 1
		}

	}

	return false
}

func binarySearchRecursive(arr []int, v int, lo int, hi int) bool {
	for lo < hi {
		m := int(lo + (hi-lo)/2)
		n := arr[m]

		if v == n {
			return true
		} else if n > v {
			return binarySearchRecursive(arr, v, lo, m)
		} else {
			return binarySearchRecursive(arr, v, m+1, hi)
		}
	}
	return false
}
