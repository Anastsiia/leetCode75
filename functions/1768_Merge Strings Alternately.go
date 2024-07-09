/*
	You are given two strings word1 and word2. Merge the strings by adding letters in alternating order, starting with word1. If a string is longer than the other, append the additional letters onto the end of the merged string.

Return the merged string.
*/
package functions

func MergeAlternately(word1 string, word2 string) string {
	res := []byte{}
	i := 0
	for ; i < len(word2) || i < len(word1); i++ {
		if i < len(word1) {
			res = append(res, word1[i])
		}
		if i < len(word2) {
			res = append(res, word2[i])
		}
	}
	return string(res)
}
