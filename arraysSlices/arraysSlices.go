package arraysslices

// SumArray adds up the numbers inside an array
func SumArray(numbers [5]int) int {

	result := 0
	for i := 0; i < 5; i++ {
		result += numbers[i]
	}

	return result
}

// SumArrayUsingRange adds up the numbers inside an array using the range function
func SumArrayUsingRange(numbers [5]int) int {
	result := 0

	for _, number := range numbers {
		result += number
	}

	return result
}

// SumSlice takes a slice of numbers are adds them up
func SumSlice(numbers []int) int {
	result := 0

	for _, number := range numbers {
		result += number
	}

	return result
}

// SumAll takes multiple slices, adds the numbers of each slice, and returns a slice with the results of each input slice
func SumAll(numbersToSum ...[]int) (sums []int) {

	lengthOfNumbers := len(numbersToSum)
	sums = make([]int, lengthOfNumbers)

	for i, numbers := range numbersToSum {
		sums[i] = SumSlice(numbers)
	}

	return
}

// SumAllRefactored takes multiple slices, adds the numbers of each slice, and returns a slice with the results of each input slice (refactored)
func SumAllRefactored(numbersToSum ...[]int) []int {

	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, SumSlice(numbers))
	}
	return sums
}

// SumAllTails takes multiple slices and returns the sum of values after the first value in each slice, in a new slice
func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, SumSlice(tail))
		}
	}

	return sums
}
