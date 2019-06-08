package reverseStr

func reverseStr(s string, k int) string {
	rs := []byte(s)

	for l := 0; l < len(rs); l += 2 * k {
		r := l + k - 1
		if len(s) < k-1 || len(s) < k+l {
			r = len(rs) - 1
		}
		rs = rev(rs, l, r)
	}
	return string(rs)
}

func rev(rs []byte, l, r int) []byte {
	for ; l < r; l, r = l+1, r-1 {
		rs[l], rs[r] = rs[r], rs[l]
	}
	return rs
}
