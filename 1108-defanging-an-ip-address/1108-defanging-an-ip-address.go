func defangIPaddr(ad string) string {
    chs :=  strings.Split(ad, ".")
    res := strings.Join(chs, "[.]")
    return res
}