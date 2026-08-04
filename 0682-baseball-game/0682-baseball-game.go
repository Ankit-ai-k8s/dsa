func calPoints(operations []string) int {
	st := []int{}

	for _, op := range operations {

		nm, err := strconv.Atoi(op)

		if err == nil {
			push(&st, nm)
		} else {
			if op == "C" {
				pop(&st)
			} else if op == "+" {
				n2, _ := pop(&st)
				n1, _ := pop(&st)
				sm := n1 + n2
				push(&st, n1)
				push(&st, n2)
				push(&st, sm)
			} else {
				st = append(st, st[len(st)-1]*2)
			}
		}
	}
	res := 0
	for _, nm := range st {
		res += nm
	}

	return res
}

func push(st *[]int, ele int) {
	*st = append(*st, ele)
}

func pop(st *[]int) (int, bool) {
	if len(*st) == 0 {
		return -1, false
	}
	last := len(*st) - 1
	ele := (*st)[last]
	*st = (*st)[:last]
	return ele, true
}