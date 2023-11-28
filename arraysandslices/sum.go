package arraysandslices

func SumArray(arr [5]int) int {
	var total int
	for _, a := range arr {
		total += a
	}

	return total
}

func SumSlice(sli []int) int {
	var total int
	for _, a := range sli {
		total += a
	}

	return total
}

func SumAll(sl1 ...[]int) []int {
	var sums []int

	for _, numbers := range sl1 {
		sums = append(sums, SumSlice(numbers))
	}

	return sums
}
