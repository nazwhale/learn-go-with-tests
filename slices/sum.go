package slices

func Sum(numbers []int) int {
	var count int
	for _, n := range numbers {
		count += n
	}
	return count
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, nums := range numbersToSum {
		sums = append(sums, Sum(nums))
	}

	return sums
}

func SumAllTails(heads int, numbersToSum ...[]int) []int {
	var sums []int
	for _, nums := range numbersToSum {
		if len(nums) < heads {
			sums = append(sums, 0)
			continue
		}
		tail := nums[heads:]
		sums = append(sums, Sum(tail))
	}

	return sums
}
