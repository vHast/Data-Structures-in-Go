package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

func maxDepth(s string) int {
	parenthesesString := charFilter(s)
	parenthesesArr := strings.Split(parenthesesString, "")
	parenthesesCount := 0 // In case we start without parentheses
	var parenthesesCountArr []int
	isInside := false

	for i := 0; i < len(parenthesesArr); i++ {
		if parenthesesArr[i] == "(" {
			parenthesesCount++
			isInside = true
		} else {
			parenthesesCountArr = append(parenthesesCountArr, parenthesesCount)

			if isInside {
				parenthesesCount = 1
			} else {
				parenthesesCount = 0
			}
		}
	}

	fmt.Println(parenthesesArr)
	sort.Ints(parenthesesCountArr)
	return parenthesesCountArr[len(parenthesesCountArr)-1]
}

func charFilter(s string) string {
	regexPattern := `[\d+*-/]`
	regCompiled := regexp.MustCompile(regexPattern)
	return regCompiled.ReplaceAllString(s, "")
}

func main() {
	fmt.Println(maxDepth("(1+(2*3)+((8)/4))+1"))
	fmt.Println(maxDepth("(1)+((2))+(((3)))"))
	fmt.Println(maxDepth("1+(2*3)/(2-1)"))
}
