/*
Given an array of characters chars, compress it using the following algorithm:

Begin with an empty string s. For each group of consecutive repeating characters in chars:

If the group's length is 1, append the character to s.
Otherwise, append the character followed by the group's length.
The compressed string s should not be returned separately, but instead, be stored in the input character array chars. Note that group lengths that are 10 or longer will be split into multiple characters in chars.

After you are done modifying the input array, return the new length of the array.

You must write an algorithm that uses only constant extra space.
*/
package functions

import (
	"strconv"
)

func Compress(chars []byte) int {
	if len(chars) == 1 || len(chars) == 0 {
		return len(chars)
	}
	res := string(chars[0])
	count := 1
	for i := 1; i < len(chars); i++ {

		if chars[i] != chars[i-1] {
			if count > 1 {
				res += strconv.Itoa(count)
			}
			res += string(chars[i])
			count = 1
		} else if chars[i] == chars[i-1] {
			count++
		}
	}

	if count > 1 {
		res += strconv.Itoa(count)
	}
	for i, v := range res {
		chars[i] = byte(v)
	}
	chars = chars[0:len(res)]
	return len(res)
}

/* second solution with one loop
func Compress(chars []byte) int {
	if len(chars) == 1 || len(chars) == 0 {
		return len(chars)
	}
	res := ""
	ch := 0
	for i := 0; i < len(chars); {
		count := 1
		j := i + 1
		for ; j < len(chars) && chars[j] == chars[j-1]; j++ {

			count++
		}
		res += string(chars[i])
		chars[ch] = chars[j-1]
		ch++
		if count > 1 {
			res += strconv.Itoa(count)
			for _, v := range strconv.Itoa(count) {
				chars[ch] = byte(v)
				ch++
			}
		}
		i = j
	}
	chars = chars[0:len(res)]
	return len(res)
}
*/
