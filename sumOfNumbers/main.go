/*
Given two integers a and b, which can be positive or negative, find the sum of all the integers between and including them and return it. If the two numbers are equal return a or b.

Note: a and b are not ordered!
Examples (a, b) --> output (explanation)

(1, 0) --> 1 (1 + 0 = 1)
(1, 2) --> 3 (1 + 2 = 3)
(0, 1) --> 1 (0 + 1 = 1)
(1, 1) --> 1 (1 since both are same)
(-1, 0) --> -1 (-1 + 0 = -1)
(-1, 2) --> 2 (-1 + 0 + 1 + 2 = 2)


*/

package main

import "fmt"

func GetSum(a, b int) int {

	var biggerNumber int
	var smallerNumber int

	if a > b {
		biggerNumber = a
		smallerNumber = b
	} else {
		biggerNumber = b
		smallerNumber = a
	}

	var arrNums int

	// Find number in-between values
	for i := smallerNumber; i <= biggerNumber; i++ {
		// fmt.Printf("Checking number %d while biggerNumber is %d\n", i, biggerNumber)
		arrNums += i
	}

	return arrNums
}

func main() {
	fmt.Println(GetSum(0, 1))
	fmt.Println(GetSum(1, 2))
	fmt.Println(GetSum(5, -1))
}
