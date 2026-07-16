func longestConsecutive(nums []int) int {

	set := map[int]struct{}{}
	for _, i := range nums {
		set[i] = struct{}{}
	}

	mx := 0

	for i,_ := range set {

		// Not the start of a sequence
		if _, ok := set[i-1]; ok {
			continue
		}

		nm := i
		cnt := 1

		for {
			if _, ok := set[nm+1]; !ok {
				break
			}
			cnt++
			nm++
		}

		if cnt > mx {
			mx = cnt
		}
	}

	return mx

}