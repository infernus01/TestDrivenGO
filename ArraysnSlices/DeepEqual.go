package ArraysnSlices

// func SumAll(numbersToSum ...[]int) []int {
// 	// lenOfNumbers := len(numbersToSum)
// 	// sums := make([]int, lenOfNumbers)

// 	// for i, numbers := range numbersToSum {
// 	// 	sums[i] = Sum(numbers)
// 	// }

// 	// var sums []int
// 	// for _, numbers := range numbersToSum {
// 	// 	sums = append(sums, Sum(numbers)) //append function takes a slice and a new value, then returns a new slice with all the items in it. 
// 	// }
// 	// return sums
// }

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _,numbers:=range numbersToSum {
		sums=append(sums,Sum(numbers[1:]))
	}
	return sums
}
