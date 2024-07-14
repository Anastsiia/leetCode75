/*
You are given an integer array nums and an integer k.

In one operation, you can pick two numbers from the array whose sum equals k and remove them from the array.

Return the maximum number of operations you can perform on the array.
*/
package functions

// map solution, o(n)
func MaxOperations(nums []int, k int) int {
	sum := 0
	numsList := map[int]int{}
	for _, v := range nums {
		if numsList[k-v] > 0 {
			numsList[k-v]--
			sum++
		} else {
			numsList[v]++
		}
	}
	return sum
}

// two pointer solution
/*func MaxOperations(nums []int, k int) int {
	sum, left, right := 0, 0, len(nums)-1
	sort.Ints(nums)

	for left < right {
		currSum := nums[left] + nums[right]
		if currSum == k {
			sum++
			left++
			right--
		} else if currSum > k {
			right--
		} else {
			left++
		}
	}
	return sum
}
*/
