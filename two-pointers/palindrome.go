package main

import (
	"fmt"
	"strings"
)

func main() {
	s := " "
	fmt.Println(isPalindrome(s))
}

func isPalindrome(s string) bool {
	s = removeNonAlphanumericCharacters(s)

	if len(s) == 0 {
		return true
	}

	for i := 0; i <= len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}

	return true
}

func removeNonAlphanumericCharacters(s string) string {
	var result strings.Builder
	for i := 0; i < len(s); i++ {
		b := s[i]
		if ('a' <= b && b <= 'z') || ('A' <= b && b <= 'Z') || ('0' <= b && b <= '9') {
			result.WriteByte(b)
		}
	}
	return strings.ToLower(result.String())
}
