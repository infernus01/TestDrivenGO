package ArraysnSlices

// Arrays are fixed length and slices are dynamic length.
// Arrays syntax: [5]int{1,2,3,4,5}
// Slices syntax: []int{1,2,3,4,5}
func Sum(numbers []int) int { //slice of ints
	sum := 0
	// for i:=0;i<5;i++{
	// 	sum+=numbers[i]
	// }
	for _, number := range numbers { //range returns the index and the value and we are not using the index so we ignore it using _.
		sum += number
	}
	return sum
}
