// https://leetcode.com/problems/valid-palindrome/
package main

import (
	"strings"
)

func isPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}
	s = strings.ToLower(s)
	head := 0
	tail := len(s) - 1
	for head < tail {
		if !isValid(s[head]) {
			head++
			continue
		}
		if !isValid(s[tail]) {
			tail--
			continue
		}
		if s[head] != s[tail] {
			return false
		}
		head++
		tail--
	}
	return true
}

func isValid(b byte) bool {
	if b >= 'a' && b <= 'z' || b >= '0' && b <= '9' {
		return true
	}
	return false
}
