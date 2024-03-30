/*

Your task is to convert strings to how they would be written by Jaden Smith. The strings are actual quotes from Jaden Smith, but they are not capitalized in the same way he originally typed them.

Example:

Not Jaden-Cased: "How can mirrors be real if our eyes aren't real"
Jaden-Cased:     "How Can Mirrors Be Real If Our Eyes Aren't Real"

*/

package main

import (
	"fmt"
	"strings"
)

func ToJadenCase(str string) string {
	var arrWords = strings.Split(str, " ")
	var capitalizedWords []string

	for i := 0; i < len(arrWords); i++ {
		wordArr := strings.Split(arrWords[i], "")
		wordArr[0] = strings.ToUpper(wordArr[0])
		finalWord := strings.Join(wordArr, "")
		capitalizedWords = append(capitalizedWords, finalWord)
	}

	formattedPhrase := strings.Join(capitalizedWords, " ")
	return formattedPhrase
}

func main() {
	fmt.Println(ToJadenCase("How can mirrors be real if our eyes aren't real"))
}
