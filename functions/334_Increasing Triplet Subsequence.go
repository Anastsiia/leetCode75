/*
Given an integer array nums, return true if there exists a triple of indices (i, j, k) such that i < j < k and nums[i] < nums[j] < nums[k]. If no such indices exists, return false.
*/
package functions

import (
	"math"
)

func IncreasingTriplet(nums []int) bool {
	first, second := math.MaxInt64, math.MaxInt64

	for i := 0; i < len(nums); i++ {
		if nums[i] <= first {
			first = nums[i]
		} else if nums[i] <= second {
			second = nums[i]
		} else {
			return true
		}
	}
	return false
}
