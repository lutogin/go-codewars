package snailSort

/**
Given an n x n array, return the array elements arranged from outermost elements to the middle element, traveling clockwise.

array = [[1,2,3],
         [4,5,6],
         [7,8,9]]
snail(array) #=> [1,2,3,6,9,8,7,4,5]
For better understanding, please follow the numbers of the next array consecutively:

array = [[1,2,3],
         [8,9,4],
         [7,6,5]]
snail(array) #=> [1,2,3,4,5,6,7,8,9]
This image will illustrate things more clearly:


NOTE: The idea is not sort the elements from the lowest value to the highest; the idea is to traverse the 2-d array in a clockwise snailshell pattern.

NOTE 2: The 0x0 (empty matrix) is represented as en empty array inside an array [[]].

https://www.codewars.com/kata/521c2db8ddc89b9b7a0000c1/train/go
*/

func Run(snailMap [][]int) []int {
	var result []int
	if len(snailMap[0]) == 0 {
		return []int{}
	}

	copySnailMap := snailMap[:]

	for true {
		result = append(result, copySnailMap[0]...) // first row always be in start
		copySnailMap = copySnailMap[1:]

		rows := len(copySnailMap)
		if rows == 0 {
			break
		}

		cols := len(copySnailMap[0])

		newMatrix := make([][]int, cols)
		for i := range newMatrix {
			newMatrix[i] = make([]int, rows)
		}

		for i := 0; i < cols; i++ {
			for j := 0; j < rows; j++ {
				newMatrix[i][j] = copySnailMap[j][cols-i-1]
			}
		}

		copySnailMap = newMatrix[:]
	}

	return result
}
