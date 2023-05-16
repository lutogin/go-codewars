package rot13

import (
	"regexp"
	"strings"
)

/**
How can you tell an extrovert from an introvert at NSA?
Va gur ryringbef, gur rkgebireg ybbxf ng gur BGURE thl'f fubrf.

I found this joke on USENET, but the punchline is scrambled. Maybe you can decipher it?
According to Wikipedia, ROT13 is frequently used to obfuscate jokes on USENET.

For this task you're only supposed to substitute characters. Not spaces, punctuation, numbers, etc.

Test examples:

"EBG13 rknzcyr." -> "ROT13 example."

"This is my first ROT13 excercise!" -> "Guvf vf zl svefg EBG13 rkprepvfr!"

https://www.codewars.com/kata/52223df9e8f98c7aa7000062/train/go
*/

func Run(msg string) string {
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	shiftedAlpabetArr := strings.Split("NOPQRSTUVWXYZABCDEFGHIJKLMnopqrstuvwxyzabcdefghijklm", "")

	input := strings.Split(msg, "")
	var result string

	for _, v := range input {
		isItWord, _ := regexp.MatchString(`[a-zA-Z]`, v)
		if !isItWord {
			result += v
			continue
		}

		idx := strings.Index(alphabet, v)

		result += shiftedAlpabetArr[idx]
	}

	return result
}
