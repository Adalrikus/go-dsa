package array

import "sort"

// ZeroSumTriplets solves the problem in O(n^2) time and O(1) space.
func ZeroSumTriplets(list []int) [][]int {
	result := make([][]int, 0)
	if len(list) < 3 {
		return result
	}
	sort.Ints(list)
	n := len(list)
	for i := 0; i < n; i++ {
		if i > 0 && list[i] == list[i-1] {
			continue
		}

		start, end := i+1, n-1
		for start < end {
			triplets := list[i] + list[start] + list[end]
			switch {
			case triplets > 0:
				end--
			case triplets < 0:
				start++
			default:
				result = append(result, []int{list[i], list[start], list[end]})
				start++
				for list[start] == list[start-1] && start < end {
					start++
				}
			}
		}
	}
	return result
}
