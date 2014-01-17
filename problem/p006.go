package problem

// Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.
func P006() int {
	var sum int
	var square int
	for i := 1; i <= 100; i++ {
		sum += i * i
		square += i
	}
	return square*square - sum
}
