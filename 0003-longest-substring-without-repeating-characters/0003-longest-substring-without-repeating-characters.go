func lengthOfLongestSubstring(s string) int {
	chars := map[byte]int{}
	start := 0
	end := 0
	res := 0
	for end < len(s) {
		ch := s[end]
		chars[ch]++

		for len(chars) < (end - start + 1) {
            chs:= s[start]
            chars[chs]--
            if chars[chs] == 0 {
                delete(chars, chs)
            }
            start++
		}
        res = max(res, end - start + 1)
        end++
	}
	return res
}