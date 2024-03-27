/*
Given a set of numbers, return the additive inverse of each. Each positive becomes negatives, and the negatives become positives.

invert([1,2,3,4,5]) == [-1,-2,-3,-4,-5]
invert([1,-2,3,-4,5]) == [-1,2,-3,4,-5]
invert([]) == []

*/

package main

func Invert(array []int) []int {
	// If the input array is empty, return it as is
	if len(array) == 0 {
		return array
	}

	var endProduct []int

	for i := 0; i < len(array); i++ {
		// Convert positive numbers to negative and vice versa
		endProduct = append(endProduct, -array[i])
	}

	return endProduct
}

func main() {
	Invert([]int{1, 2, 3, 4, 5, -6, -2})
}
