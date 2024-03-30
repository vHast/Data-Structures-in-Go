/*
Build a function that returns an array of integers from n to 1 where n>0.

Example : n=5 --> [5,4,3,2,1]

*/

package main

import "fmt"

func ReverseSeq(n int) []int {
	var numbers []int

	for i := n; i >= 1; i-- {
		numbers = append(numbers, i)
	}

	return numbers
}

func main() {
	fmt.Println(ReverseSeq(5))
}
