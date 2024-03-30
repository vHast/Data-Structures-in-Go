/*
Given an array of strings strs, group the anagrams together. You can return the answer in any order.

An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.



Example 1:

Input: strs = ["eat","tea","tan","ate","nat","bat"]
Output: [["bat"],["nat","tan"],["ate","eat","tea"]]

*/

package main

import (
	"fmt"
	"sort"
	"strings"
)

func groupAnagrams(strs []string) []string {
	// var output [][]string
	var sortedStrings []string

	for i := 0; len(strs) > i; i++ {
		outputString := splitter(strs[i])
		sortedStrings = append(sortedStrings, outputString)
	}

	return sortedStrings
}

func splitter(str string) string {
	splitStr := strings.Split(str, "")
	sort.Strings(splitStr)
	return strings.Join(splitStr, "")
}

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}
