func characterReplacement(s string, k int) int {
    m:= [26]int{}
	start := 0
	end := 0
	res := 0
	mx := 0
	for end < len(s) {
		ch := s[end]
		m[ch - 'A']++
		mx = max(m[ch - 'A'], mx)
		for ((end - start + 1) - mx) > k {
			chs := s[start]
			m[chs - 'A']--
			start++
		}
		res = max(res, end - start + 1)
		end++
	}

	return res
}