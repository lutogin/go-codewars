package diagonalSum

/**
Given a square matrix mat, return the sum of the matrix diagonals.

Only include the sum of all the elements on the primary diagonal and all the elements on the secondary diagonal that are not part of the primary diagonal.

Input: mat = [[1,2,3],
              [4,5,6],
              [7,8,9]]
Output: 25
Explanation: Diagonals sum: 1 + 5 + 9 + 3 + 7 = 25
Notice that element mat[1][1] = 5 is counted only once.

Input: mat = [[1,1,1,1],
              [1,1,1,1],
              [1,1,1,1],
              [1,1,1,1]]
Output: 8

Constraints:

n == mat.length == mat[i].length
1 <= n <= 100
1 <= mat[i][j] <= 100

https://leetcode.com/problems/matrix-diagonal-sum/
*/

func GetSample() [][]int {
	return [][]int{
		{1, 1, 1, 1},
		{1, 1, 1, 1},
		{1, 1, 1, 1},
		{1, 1, 1, 1},
	}
}

func GetSample2() [][]int {
	return [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
}

func GetSample3() [][]int {
	return [][]int{
		{5},
	}
}

func Run(mat [][]int) int {
	length := len(mat)
	if length != len(mat[0]) || mat[0][0] > 100 || mat[0][0] < 1 {
		return 0
	}

	if length == 1 {
		return mat[0][0]
	}

	//middle := int(math.Round(float64(len(mat) / 2)))
	sum := 0
	for i, v := range mat {
		if i == length-1-i {
			sum += v[i]
		} else {
			sum += v[i] + v[length-1-i]
		}
	}

	return sum
}
