package three_sum

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	results := make([][]int, 0)
	for i := 0; i < len(nums)-2; i++ {
		left := i + 1
		right := len(nums) - 1

		if nums[i] > 0 {
			break
		}

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				results = append(results, []int{nums[i], nums[left], nums[right]})
				left++
				right--
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return results
}
