/*
Given a binary array nums and an integer k, return the maximum number of consecutive 1's in the array if you can flip at most k 0's.
*/
package functions

// sliding window, o(n)
func LongestOnes(nums []int, k int) int {
	res := 0
	currSum, start := 0, 0
	for end := 0; end < len(nums); end++ {
		currSum += nums[end]
		if (end-start+1)-currSum <= k {
			if res < end-start+1 {
				res = end - start + 1
			}
		} else {
			currSum -= nums[start]
			start++
		}
	}
	return res
}

//intuitive solution O(n^2)
/*func LongestOnes(nums []int, k int) int {
	zeros, count, res := k, 0, 0
	for i := 0; i < len(nums); i++ {

		for a := i; a < len(nums); a++ {
			if nums[a] == 1 {
				count++
			} else if nums[a] == 0 {
				if zeros > 0 {
					zeros--
					count++
				} else {
					break
				}
			}
		}
		if count > res {
			res = count
		}
		count = 0
		zeros = k
	}
	return res
}
*/
