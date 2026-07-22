func maxArea(hg []int) int {
	start := 0
	end := len(hg) - 1
	res := 0
	for start < end {
		mn := 0
		if hg[start] > hg[end] {
			mn = hg[end]
		} else {
			mn = hg[start]
		}
		cur := mn * (end - start)

		if cur > res {
			res = cur
		}

		if hg[start] > hg[end] {
			end--
		} else {
			start++
		}
	}
	return res
}