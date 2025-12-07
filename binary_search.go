package binary_search

// LowerBound finds the position of the first occurrence
// not less than (either equal or greater) the given value.
func LowerBound(arr []int, target int) (int, bool) {
	left, right := -1, len(arr)-1

	for right-left > 1 {
		mid := (left + right) >> 1
		if arr[mid] < target {
			left = mid
		} else {
			right = mid
		}
	}

	if arr[right] < target {
		return -1, false
	}

	return right, true
}
