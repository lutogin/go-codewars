package findTheNextPerfectSquare

import "math"

/**
You might know some pretty large perfect squares. But what about the NEXT one?

Complete the findNextSquare method that finds the next integral perfect square after the one passed as a parameter. Recall that an integral perfect square is an integer n such that sqrt(n) is also an integer.

If the parameter is itself not a perfect square then -1 should be returned. You may assume the parameter is non-negative.

Examples:(Input --> Output)

121 --> 144
625 --> 676
114 --> -1 since 114 is not a perfect square

https://www.codewars.com/kata/56269eb78ad2e4ced1000013/train/go
*/

func Run(sq int64) int64 {
	s := math.Sqrt(float64(sq))
	if math.Mod(s, 1) != 0 {
		return -1
	}

	return int64(math.Pow(s+1, 2))
}
