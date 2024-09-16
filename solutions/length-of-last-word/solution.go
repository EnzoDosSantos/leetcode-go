package solutions

import "strings"

// PROBLEM URL: https://leetcode.com/problems/length-of-last-word/

func LengthOfLastWord(s string) int {
	splitedString := strings.Split(strings.TrimSpace(s), " ")

	lastWord := splitedString[len(splitedString)-1]

	return len(lastWord)
}
