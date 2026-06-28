func categorizeBox(len int, wi int, hg int, ms int) string {

	isH := isHeavy(ms)
	isB := isBulky(len, wi, hg)
	if isH && isB {
		return "Both"
	} else if isH {
		return "Heavy"
	} else if isB {
		return "Bulky"
	}
	return "Neither"
}

func isBulky(len int, wi int, hg int) bool {
	bdm := 10000
	if wi >= bdm || hg >= bdm || len >= bdm {
		return true
	}
	vol := wi * hg * len
	if vol >= 1000000000 {
		return true
	}
	return false
}

func isHeavy(ms int) bool {
	return ms >= 100
}