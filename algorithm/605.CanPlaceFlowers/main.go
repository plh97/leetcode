package canPlaceFlowers

func canPlaceFlowers(flowerbed []int, n int) bool {
	for i := 0; i < len(flowerbed); i++ {
		// 检测左边,大于0才检测,
		if i > 0 && flowerbed[i-1] == 1 {
			continue
		}
		// 检测当前
		if flowerbed[i] == 1 {
			continue
		}
		// 检测右边
		if i+1 < len(flowerbed) && flowerbed[i+1] == 1 {
			continue
		}
		n--
		flowerbed[i] = 1
	}
	return n < 1
}
