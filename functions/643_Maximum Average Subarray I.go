/*
You are given an integer array nums consisting of n elements, and an integer k.

Find a contiguous subarray whose length is equal to k that has the maximum average value and return this value. Any answer with a calculation error less than 10-5 will be accepted.
*/
package functions

func FindMaxAverage(nums []int, k int) float64 {
	if k == 0 || k > len(nums) {
		return 0.0
	}
	res := 0
	left := 0
	for ; left <= k-1; left++ {
		res += nums[left]
	}
	left = 0
	currSum := res
	for ; left+k <= len(nums)-1; left++ {
		currSum += nums[left+k] - nums[left]
		if currSum > res {
			res = currSum
		}
	}
	return float64(res) / float64(k)
}
