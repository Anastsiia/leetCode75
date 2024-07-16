/*
Given a string s, reverse only all the vowels in the string and return it.

The vowels are 'a', 'e', 'i', 'o', and 'u', and they can appear in both lower and upper cases, more than once.
*/
package functions

func IsVowel(c rune) bool {
	check := "aAeEiIoOuU"
	for _, v := range check {
		if v == c {
			return true
		}
	}
	return false
}

func ReverseVowels(s string) string {
	res := []rune(s)
	right, left := len(s)-1, 0
	for left < right {
		for ; left < right && !IsVowel(res[left]); left++ {
		}
		for ; right > left && !IsVowel(res[right]); right-- {
		}
		res[left], res[right] = res[right], res[left]
		left++
		right--
	}
	return string(res)
}

//straight forward solution o(n)
/*
func ReverseVowels(s string) string {
	vowels := ""
	for _, v := range s {
		if isVowel(v) {
			vowels += string(v)
		}
	}
	if len(vowels) == 0 {
		return s
	}
	i := len(vowels) - 1
	res := ""
	for _, v := range s {
		if isVowel(v) {
			res += string(vowels[i])
			i--
		} else {
			res += string(v)
		}
	}
	return res
}*/
