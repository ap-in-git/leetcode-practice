package main

func isPalindrome(s string) bool {
	var newStr []rune
	for _, el := range s {
		if !(el >= '0' && el <= '9' || el >= 'a' && el <= 'z' || el >= 'A' && el <= 'Z') {
			continue
		}
		if el >= 'a' && el <= 'z' {
			el = el - 32
		}
		newStr = append(newStr, el)
	}
	l := 0
	r := len(newStr) - 1
	for l <= r {
		if newStr[l] != newStr[r] {
			return false
		}
		l++
		r--
	}

	return true
}
