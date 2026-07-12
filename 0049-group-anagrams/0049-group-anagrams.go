func groupAnagrams(strs []string) [][]string {
	m := map[string][]string{}

	for _, s := range strs {
		ln := len(s)
		byt := make([]byte, 26)

		for i := 0; i < ln; i++ {
			byt[int(s[i]) - 'a']++
		}
		 ls:= m[string(byt)]
         ls = append(ls, s)
         m[string(byt)] = ls
	}

	// fmt.Println(m)

    res:= [][]string{}

    for _,v := range m {
        res = append(res, v)
    } 

	return res
}