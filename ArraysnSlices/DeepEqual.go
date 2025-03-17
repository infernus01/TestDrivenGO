package ArraysnSlices

func SumAll(numbersToSum ...[]int) []int {
	lenOfNumbers := len(numbersToSum)
	sums := make([]int, lenOfNumbers)

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}
	return sums
}

