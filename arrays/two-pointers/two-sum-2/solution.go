package two_sum_2

func TwoSumIINaive(input []int, target int) []int {

	for i, _ := range input {
		lo := i
		for hi := lo + 1; hi < len(input); hi++ {
			if (input[lo] + input[hi]) == target {
				return []int{lo + 1, hi + 1}
			}
		}
	}

	return []int{}
}

func TwoSumII(input []int, target int) []int {
	lo := 0
	hi := len(input) - 1
	for lo < hi {
		sum := input[lo] + input[hi]
		if sum == target {
			return []int{lo + 1, hi + 1}
		} else if sum < target {
			lo++
		} else {
			hi--
		}
	}
	return []int{}
}
