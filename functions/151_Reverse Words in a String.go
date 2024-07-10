/*
Given an input string s, reverse the order of the words.

A word is defined as a sequence of non-space characters. The words in s will be separated by at least one space.

Return a string of the words in reverse order concatenated by a single space.

Note that s may contain leading or trailing spaces or multiple spaces between two words. The returned string should only have a single space separating the words. Do not include any extra spaces.
*/
package functions

// straight forward solution
func ReverseWords(s string) string {
	arr := []string{}
	temp := ""
	for _, v := range s {

		if v == ' ' && temp != "" {
			arr = append(arr, temp)
			temp = ""
		} else if v == ' ' && temp == "" {
			continue
		} else {
			temp += string(v)
		}
	}
	if temp != "" {
		arr = append(arr, temp)
		temp = ""
	}
	for i := len(arr) - 1; i >= 0; i-- {
		temp += arr[i] + " "
	}
	return temp[:len(temp)-1]
}

//solution with strings.Fields
/*func ReverseWords(s string) string {
	sArr := strings.Fields(s)
	res := ""
	for i := len(sArr) - 1; i >= 0; i-- {
		res += sArr[i] + " "

	}
	res = strings.Trim(res, " ")
	return res
}
*/
