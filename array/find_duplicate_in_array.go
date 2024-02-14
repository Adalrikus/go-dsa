package array

// FindDuplicate solves the problem in O(n) time and O(1) space.
type set map[int]struct{}

func (s set) add(number int) {
	s[number] = struct{}{}
}

func (s set) has(number int) bool {
	_, ok := s[number]
	return ok
}

func FindDuplicate(list []int) int {
	numSet := set{}
	for _, v := range list {
		if numSet.has(v) {
			return v
		}
		numSet.add(v)
	}
	return -1
}
