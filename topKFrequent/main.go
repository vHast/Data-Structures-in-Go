/*
Given an integer array nums and an integer k, return the k most frequent elements. You may return the answer in any order.



Example 1:

Input: nums = [1,1,1,2,2,3], k = 2
Output: [1,2]

Example 2:

Input: nums = [1], k = 1
Output: [1]


*/

package main

import "fmt"

func topKFrequent(nums []int, k int) []int {

	var values []int

	for i := 0; i < k; i++ {
		valueToAdd := 1
		for j := 0; j < len(nums)-1; j++ {

			if nums[j] == nums[j]+1 {
				valueToAdd += 1
			}
		}
		values = append(values, valueToAdd)
	}
	return values
}

func main() {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
}
