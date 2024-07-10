/*
You have a long flowerbed in which some of the plots are planted, and some are not. However, flowers cannot be planted in adjacent plots.

Given an integer array flowerbed containing 0's and 1's, where 0 means empty and 1 means not empty, and an integer n, return true if n new flowers can be planted in the flowerbed without violating the no-adjacent-flowers rule and false otherwise.
*/
package functions

import "fmt"

func sumArr(arr []int) int {
	res := 0
	for _, v := range arr {
		res += v
	}
	return res
}

func CanPlaceFlowers(flowerbed []int, n int) bool {
	// if n == 1 && len(flowerbed) == 1 {
	// 	return flowerbed[0] == 0
	// }
	// if (len(flowerbed)-sumArr(flowerbed))/2 < n {
	// 	return false
	// }
	for i := 0; i < len(flowerbed) && n != 0; i++ {
		if flowerbed[i] == 0 && (i == 0 || flowerbed[i-1] == 0) && (i == len(flowerbed)-1 || flowerbed[i+1] == 0) {
			flowerbed[i] = 1
			n--
			i++
			fmt.Println(flowerbed)
		}
	}
	return n == 0
}

// 0, 0, 0,
// 2
