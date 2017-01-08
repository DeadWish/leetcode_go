func reverseString(s string) string {
	st := []byte(s)
	slen := len(st)
	for s, e := 0, slen-1; s < e; s, e = s+1, e-1 {
		st[s], st[e] = st[e], st[s]
	}
	return string(st)
}