/*

Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.

Example 1:

Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

*/

package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var output []int

	for i := 0; len(nums) > i; i++ {
		for j := 1 + i; len(nums) > j; j++ {
			if nums[i]+nums[j] == target {
				output = append(output, i, j)
			}
		}
	}

	return output
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println(twoSum([]int{3, 3}, 6))
}
