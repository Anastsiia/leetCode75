/*
There are n kids with candies. You are given an integer array candies, where each candies[i] represents the number of candies the ith kid has, and an integer extraCandies, denoting the number of extra candies that you have.

Return a boolean array result of length n, where result[i] is true if, after giving the ith kid all the extraCandies, they will have the greatest number of candies among all the kids, or false otherwise.

Note that multiple kids can have the greatest number of candies.
*/
package functions

func arrayMax(arr []int) int {
	res := arr[0]
	for _, v := range arr {
		if v > res {
			res = v
		}
	}
	return res
}

func KidsWithCandies(candies []int, extraCandies int) []bool {
	res := make([]bool, len(candies))
	max := arrayMax(candies)
	for i := 0; i < len(candies); i++ {
		res[i] = candies[i]+extraCandies >= max
	}
	return res
}
