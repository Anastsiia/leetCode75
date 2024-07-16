/*
Given a string s and an integer k, return the maximum number of vowel letters in any substring of s with length k.

Vowel letters in English are 'a', 'e', 'i', 'o', and 'u'.
*/
package functions

func MaxVowels(s string, k int) int {
	res, count := 0, 0
	for _, v := range s[:k] {
		if IsVowel(v) {
			res++
		}
	}

	count = res
	for i := 1; i+k <= len(s); i++ {
		if res >= k {
			return res
		}
		if IsVowel(rune(s[i-1])) {
			count--
		}
		if IsVowel(rune(s[i+k-1])) {
			count++
		}
		if count > res {
			res = count
		}
	}
	return res
}

//Straight forward, too slow
/*func MaxVowels(s string, k int) int {
	res, counter := 0, 0
	for i := 0; i+k <= len(s); i++ {
		counter = 0
		temp := s[i:(i + k)]
		for _, v := range temp {
			if IsVowel(v) {
				counter++
			}
			if counter >= k {
				return k
			}
		}
		if counter > res {
			res = counter
		}
	}
	return res
}*/
