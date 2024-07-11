package main

import (
	"fmt"
	"leetCode75/functions"
)

func main() {
	//1768:
	// word1 := "abc"
	// word2 := "vbn"
	// fmt.Println(functions.MergeAlternately(word1, word2))

	//1071:
	// word1 := "aab"
	// word2 := "aabaabaab"
	// fmt.Println(functions.GcdOfStrings(word1, word2))

	//1431:
	// candies := []int{4, 2, 1, 1, 2}
	// extraCandies := 0
	// fmt.Println(functions.KidsWithCandies(candies, extraCandies))

	//605:
	// flowerbed := []int{0, 0, 0}
	// n := 2
	// fmt.Println(functions.CanPlaceFlowers(flowerbed, n))

	//345:
	// str := "hll"
	// fmt.Println(functions.ReverseVowels(str))

	// 151:
	// str := "     the     sky is  blue"
	// fmt.Println(functions.ReverseWords(str))

	// 238:
	// num := []int{1, 2, 3, 4}
	// fmt.Println(functions.ProductExceptSelf(num))

	//334:
	// num := []int{0, 4, 2, 1, 0, -1, -3}
	// fmt.Println(functions.IncreasingTriplet(num))

	//443:
	// chars := []byte{'a', 'a', 'b', 'c'}
	// fmt.Println(functions.Compress(chars))

	//283:
	// nums := []int{0, 4, 2, 1, 0, -1, -3}
	// functions.MoveZeroes(nums)
	// fmt.Println(nums)

	//392:
	str1 := "baaa"
	str2 := "bbaaa"
	fmt.Println(functions.IsSubsequence(str1, str2))
}
