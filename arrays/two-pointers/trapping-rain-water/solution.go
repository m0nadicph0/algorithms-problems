package trapping_rain_water

func TrapWater(heights []int) int {
	left := 0
	right := len(heights) - 1
	leftMax := 0
	rightMax := 0
	totalWater := 0
	for left < right {
		if heights[left] <= heights[right] {
			if heights[left] > leftMax {
				leftMax = heights[left]
			} else {
				totalWater += leftMax - heights[left]
			}
			left++
		} else {
			if heights[right] > rightMax {
				rightMax = heights[right]
			} else {
				totalWater += rightMax - heights[right]
			}
			right--
		}

	}

	return totalWater
}
