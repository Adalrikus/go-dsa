package array

func Sum(list []int) int {
	var sum int
	for _, v := range list {
		sum += v
	}
	return sum
}

func EqualSubArrays(list []int) [][]int {
	sum := make([][]int, 0)
	for i := 1; i < len(list); i++ {
		subArraySum1 := Sum(list[:i])
		subArraySum2 := Sum(list[i:])
		if subArraySum1 == subArraySum2 {
			sum = append(sum, list[:i])
			sum = append(sum, list[i:])
			break
		}
	}
	return sum
}
