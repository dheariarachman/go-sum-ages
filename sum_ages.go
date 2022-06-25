package go_sum_ages

func SumAges(ages []int) int {
	result := 0

	for _, v := range ages {
		result += v
	}

	return result
}
