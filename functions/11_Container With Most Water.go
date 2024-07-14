/*
You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).

Find two lines that together with the x-axis form a container, such that the container contains the most water.

Return the maximum amount of water a container can store.

Notice that you may not slant the container.
*/
package functions

func MaxArea(height []int) int {
	res, currArea := 0, 0
	left, right := 0, len(height)-1

	for left < right {
		if height[left] > height[right] {
			currArea = (right - left) * height[right]
			right--
		} else {
			currArea = (right - left) * height[left]
			left++
		}
		if res < currArea {
			res = currArea
		}
	}
	return res
}

// straight forward, too long, time limit exceeded, o(n^2)
/*func MaxArea(height []int) int {
	res := 0
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			fmt.Println(res, i, j)
			currArea := int(math.Min(float64(height[i]), float64(height[j]))) * (j - i)
			if currArea > res {
				res = currArea
			}
		}
	}
	return res
}*/
