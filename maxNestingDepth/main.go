// Given a VPS represented as string s, return the nesting depth of s.

package main

import (
	"fmt"
	"regexp"
	"strings"
)

func maxDepth(s string) int {
	parenthesesString := charFilter(s)
	parenthesesArr := strings.Split(parenthesesString, "")
	parenthesesCount := 0 // In case we start without parentheses
	maxDepth := 0

	for i := 0; i < len(parenthesesArr); i++ {
		if parenthesesArr[i] == "(" {
			parenthesesCount++
			if parenthesesCount > maxDepth {
				maxDepth = parenthesesCount
			}
		} else if parenthesesArr[i] == ")" {
			parenthesesCount--
		}
	}

	return maxDepth
}

func charFilter(s string) string {
	regexPattern := `[\d+*-/]`
	regCompiled := regexp.MustCompile(regexPattern)
	return regCompiled.ReplaceAllString(s, "")
}

func main() {
	fmt.Println(maxDepth("(1+(2*3)+((8)/4))+1")) // Expected 3
	fmt.Println(maxDepth("(1)+((2))+(((3)))"))   // Expected 3
	fmt.Println(maxDepth("1+(2*3)/(2-1)"))       // Expected 1
}
