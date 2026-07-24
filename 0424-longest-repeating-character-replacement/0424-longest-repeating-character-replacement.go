func characterReplacement(s string, k int) int {
	m := map[string]int{}
	start := 0
	end := 0
	res := 0
	mx := 0
	for end < len(s) {
		ch := string(s[end])
		m[ch]++
		mx = max(m[ch], mx)
		for ((end - start + 1) - mx) > k {
			chs := string(s[start])
			m[chs]--
			start++
		}
		res = end - start + 1
		end++
	}

	return res
}