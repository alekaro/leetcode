// Implement strStr().

// Given two strings needle and haystack, return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.

// Clarification:

// What should we return when needle is an empty string? This is a great question to ask during an interview.
// For the purpose of this problem, we will return 0 when needle is an empty string. This is consistent to C's strstr() and Java's indexOf().

// Example 1:

// Input: haystack = "hello", needle = "ll"
// Output: 2
// Example 2:

// Input: haystack = "aaaaa", needle = "bba"
// Output: -1

// Constraints:

// 1 <= haystack.length, needle.length <= 104
// haystack and needle consist of only lowercase English characters.

package main

import "fmt"

func strStr(haystack string, needle string) int {
	var helper int
	found := true
	for i := 0; i <= len(haystack)-1; i++ {
		if haystack[i] == needle[0] && len(haystack[i:]) >= len(needle) {
			found = true
			for helper = 1; helper < len(needle); helper++ {
				if haystack[i+helper] != needle[helper] {
					found = false
					break
				}
			}
			if found {
				return i
			}
		}
	}
	return -1
}

// A little faster
func strStr2(haystack string, needle string) int {
	var helper int
	found := true
	for i := 0; i < len(haystack)+1-len(needle); i++ {
		if haystack[i] == needle[0] {
			found = true
			for helper = 1; helper < len(needle); helper++ {
				if haystack[i+helper] != needle[helper] {
					found = false
					break
				}
			}
			if found {
				return i
			}
		}
	}
	return -1
}

func main() {
	haystack := "Hello world!"
	needle := "llo"
	for i := 0; i < 5; i++ {

	}
	fmt.Println(strStr(haystack, needle))
}
