package numPairsDivisibleBy60

func numPairsDivisibleBy60(time []int) int {
	res := 0
	for i := 0; i < len(time)-1; i++ {
		for j := i + 1; j < len(time); j++ {
			if (time[i]+time[j])%60 == 0 {
				res++
			}
		}
	}
	return res
}

// func numPairsDivisibleBy60(time []int) int {
//     remainders := make([]int, 60)
    
//     var res int
    
//     for _, t := range time {
//         r := t % 60
//         var relaT int 
//         if r > 0 {
//             relaT = 60 - r
//         }
//         res += remainders[relaT]
//         remainders[r]++
//     }
    
//     return res
// }