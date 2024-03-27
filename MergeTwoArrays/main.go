package main

import (
	"fmt"
	"sort"
)

func mergeArrays(arr1, arr2 []int) []int {
	mergedArray := append(arr1, arr2...)
	sort.Ints(mergedArray)

	for i := 0; len(mergedArray)-1 > i; i++ {
		if mergedArray[i] == mergedArray[i+1] {
			mergedArray = append(mergedArray[:i], mergedArray[i+1:]...)
		}
	}

	return mergedArray
}

func main() {
	fmt.Println(mergeArrays([]int{1, 2, 3, 4}, []int{5, 4, 3, 2, 1}))
}
