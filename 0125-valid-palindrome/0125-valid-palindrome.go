func isPalindrome(s string) bool {
	res := regexp.MustCompile("[^a-zA-Z0-9]+")
	s = res.ReplaceAllString(s, "")
	s = strings.ToLower(s)
	n := len(s)
	for i := 0; i < n/2; i++ {
		if s[i] != s[n-1-i] {
			return false
		}
	}

	return true
}