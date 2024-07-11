/*
Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.

Note that you must do this in-place without making a copy of the array.
*/
package functions

// O(n) O(1) solution more elegant
func MoveZeroes(nums []int) {
	lastNum := 0
	for i, v := range nums {
		if v != 0 {
			nums[lastNum], nums[i] = nums[i], nums[lastNum]
			lastNum++
		}
	}

}

// straight forward colution o(n) o(1)
/*func MoveZeroes(nums []int) {
	for i, _ := range nums {
		fmt.Println(nums[i])
		if nums[i] == 0 && i < len(nums) {
			for j := i + 1; j < len(nums); j++ {
				if nums[j] != 0 {
					nums[i], nums[j] = nums[j], nums[i]
					break
				}
			}
		}
	}
}*/
