package maximumSubarraySum

/**
The maximum sum subarray problem consists in finding the maximum sum of a contiguous subsequence in an array or list of integers:

maxSequence [-2, 1, -3, 4, -1, 2, 1, -5, 4]
-- should be 6: [4, -1, 2, 1]
Easy case is when the list is made up of only positive numbers and the maximum sum is the sum of the whole array. If the list is made up of only negative numbers, return 0 instead.

Empty list is considered to have zero greatest sum. Note that the empty list or array is also a valid sublist/subarray.

https://www.codewars.com/kata/54521e9ec8e60bc4de000d6c/train/go
*/

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func MaximumSubarraySum(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}

	maxSoFar := numbers[0]
	maxEndingHere := numbers[0]

	for _, value := range numbers[1:] {
		maxEndingHere = maxInt(value, maxEndingHere+value)
		maxSoFar = maxInt(maxSoFar, maxEndingHere)
	}

	// if maxSoFar is negative, it means all numbers are negative in array, return 0
	if maxSoFar < 0 {
		return 0
	}

	return maxSoFar
}

func GetTestedData() []int {
	return []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
}
