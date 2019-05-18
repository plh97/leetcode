package bitwiseComplement

func bitwiseComplement(N int) int {
	x := N
	x |= 1
	x |= x >> 1
	x |= x >> 2
	x |= x >> 4
	x |= x >> 8
	return x ^ N
}

// func bitwiseComplement(N int) int {
// 	x := N
// 	x |= 1
// 	for k := x; k > 0; k /= 2 {
// 		x |= k
// 	}
// 	return x ^ N
// }
