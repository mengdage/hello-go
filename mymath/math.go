/*
Package mymath implements basic math functions.
*/
package mymath

// Add returns the sum of input integers.
func Add(num ...int) int {
	sum := 0

	for _, v := range num {
		sum += v
	}

	return sum
}
