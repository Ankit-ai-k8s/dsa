func longestCommonPrefix(strs []string) string {
	slices.SortFunc(strs, func(a, b string) int { return len(a) - len(b) })
	res := strs[0]
	for i := 1; i < len(strs); i++ {
		s2 := strs[i]
		cur := ""
		n := len(res)
		for j := 0; j < n; j++ {
			if res[j] == s2[j] {
				cur = cur + string(res[j])
			}else {
                break
            }
		}
		res = cur
	}
	return res
}