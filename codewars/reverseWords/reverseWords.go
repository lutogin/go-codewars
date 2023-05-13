package reverseWords

import "strings"

/*
Complete the function that accepts a string parameter, and reverses each word in the string. All spaces in the string should be retained.

Examples
"This is an example!" ==> "sihT si na !elpmaxe"
"double  spaces"      ==> "elbuod  secaps"

https://www.codewars.com/kata/5259b20d6021e9e14c0010d4/train/go
*/

func Run(str string) string {
	arr := strings.Split(str, " ")
	var result []string
	for _, v := range arr {
		runes := []rune(v)
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		result = append(result, string(runes))
	}

	return strings.Join(result, " ")
}
