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
)

func groupAnagrams(strs []string) [][]string {
	anagramMap := make(map[string][]string)

	for _, str := range strs {
		sorted := sortWord(str)
		anagramMap[sorted] = append(anagramMap[sorted], str)
	}

	var groupedAnagrams [][]string
	for _, anagrams := range anagramMap {
		groupedAnagrams = append(groupedAnagrams, anagrams)
	}

	return groupedAnagrams
}

func sortWord(str string) string {
	runes := []rune(str)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}
