/*
For two strings s and t, we say "t divides s" if and only if s = t + t + t + ... + t + t (i.e., t is concatenated with itself one or more times).

Given two strings str1 and str2, return the largest string x such that x divides both str1 and str2.
*/
package functions

/*func strStr(str1 string, str2 string) bool {
	if len(str2)%len(str1) != 0 {
		return false
	}
	for i := 0; i < len(str2); i += len(str1) {
		if str2[i:(i+len(str1))] != str1 {
			return false
		}
	}
	return true
}

func GcdOfStrings(str1 string, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}
	if len(str1) < len(str2) {
		str1, str2 = str2, str1
	}
	for lenRes, i := len(str2), 1; lenRes > 0 && i <= len(str2); lenRes = len(str2) / i {
		res := str2[0:lenRes]
		if strStr(res, str1) && strStr(res, str2) {
			return res
		}
		if lenRes/2 == 0 {
			break
		}
		i++
	}
	return ""
}*/

// More elegant math solution

func maxDiv(num1 int, num2 int) int {
	for i := num1; i > 0; i-- {
		if num1%i == 0 && num2%i == 0 {
			return i
		}
	}
	return 0
}

func GcdOfStrings(str1 string, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}
	return str1[0:maxDiv(len(str1), len(str2))]
}
