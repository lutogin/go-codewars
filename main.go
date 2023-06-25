package main

import (
	"codewars/codewars/josephus"
	"fmt"
)

func main() {
	fmt.Println(josephus.Josephus([]interface{}{1, 2, 3, 4, 5, 6, 7}, 3))
}
