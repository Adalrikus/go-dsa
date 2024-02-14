package array

// ProductOfAllOtherElements solves the problem in O(n) time and O(1) space.
func ProductOfAllOtherElements(list []int) []int {
	n := len(list)
	result := make([]int, n)

	prefix := 1
	for i := 0; i < n; i++ {
		result[i] = prefix
		prefix *= list[i]
	}

	suffix := 1
	for i := n - 1; i > -1; i-- {
		result[i] *= suffix
		suffix *= list[i]
	}

	return result
}
