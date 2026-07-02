func firstUniqChar(s string) int {
	m := map[rune][2]int{}
	mn := math.MaxInt
	for i, ch := range s {
		v := m[ch]
		v[0]++
		v[1] = i
		m[ch] = v
	}
	// fmt.Println(m)
	for _, v := range m {
		if v[0] == 1 && mn > v[1] {
			mn = v[1]
		}
	}
	if mn == math.MaxInt {
		return -1
	}
	return mn
}