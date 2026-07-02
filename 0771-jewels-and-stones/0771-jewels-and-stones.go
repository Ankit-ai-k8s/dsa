func numJewelsInStones(jewels string, stones string) int {
	jws := strings.Split(jewels, "")
	cnt := 0

	for _, s := range strings.Split(stones, "") {
        if slices.Contains(jws, s) {
            cnt++
        }
	}
	return cnt
}