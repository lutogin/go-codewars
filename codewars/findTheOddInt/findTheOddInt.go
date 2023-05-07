package findTheOddInt

/**
Given an array of integers, find the one that appears an odd number of times.

There will always be only one integer that appears an odd number of times.

Examples
[7] should return 7, because it occurs 1 time (which is odd).
[0] should return 0, because it occurs 1 time (which is odd).
[1,1,2] should return 2, because it occurs 1 time (which is odd).
[0,1,0,1,0] should return 0, because it occurs 3 times (which is odd).
[1,2,2,3,3,3,4,3,3,3,2,2,1] should return 4, because it appears 1 time (which is odd).

https://www.codewars.com/kata/54da5a58ea159efa38000836/train/go
*/

func Run(seq []int) int {
	if len(seq) == 1 {
		return seq[0]
	}

	m := map[int]int{}

	for _, v := range seq {
		if _, ok := m[v]; ok {
			m[v] = m[v] + 1
		} else {
			m[v] = 1
		}
	}

	for key, v2 := range m {
		if v2%2 > 0 {
			return key
		}
	}

	return 0
}
