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
	candies := []int{4, 2, 1, 1, 2}
	extraCandies := 0
	fmt.Println(functions.KidsWithCandies(candies, extraCandies))
}
