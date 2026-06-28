func defangIPaddr(ad string) string {
	chs := []byte(ad)
	res := ""
	n := len(ad)

	for i := 0; i < n; i++ {
		ch := chs[i]

		if ch == '.' {
			res = res + "[.]"
		} else {
			res = res + string(ch)
		}
	}
    
	return res
}