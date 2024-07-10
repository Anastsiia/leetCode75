/*
Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].

The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.

You must write an algorithm that runs in O(n) time and without using the division operation.
*/
package functions

func ProductExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	left, right := 1, 1
	for i, _ := range nums {
		res[i] = left
		left *= nums[i]
	}
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] *= right
		right *= nums[i]
	}
	return res
}
