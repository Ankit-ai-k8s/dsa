func containsDuplicate(nums []int) bool {
	m := map[int]bool{}

	for _, ele := range nums {
		_, ok := m[ele]
		if ok {
			return true
		}
		m[ele] = true
	}

	return false
}