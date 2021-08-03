//  strings_test.go
//
//
//  Created by JohnnyB0Y on 2021/07/27.
//  Copyright © 2021 JohnnyB0Y. All rights reserved.

package mystrings

import "testing"

func TestFirstUniqChar(t *testing.T) {
	t.Log(firstUniqChar("leetcode"))
	t.Log(firstUniqChar2("leetcode"))
	t.Log(firstUniqChar2("loveleetcode"))
	t.Log(firstUniqChar2(""))
	t.Log(firstUniqChar2("v"))
}

func TestIsPalindrome(t *testing.T) {
	t.Log(isPalindrome("A man, a plan, a canal: Panama"))
	t.Log(isPalindrome(":::::"))
	t.Log(isPalindrome("race a car"))

}

func TestStrStr(t *testing.T) {
	t.Log("--------------------------------------------")
	t.Log(strStr("hello", "ll"), "ll", 2)
	t.Log(strStr("aaaaa", "bba"), "bba", -1)
	t.Log(strStr("", ""), 0)
	t.Log(strStr("mississippi", "issip"), "issip", 4)
	t.Log(strStr("mississippi", "pi"), "pi", 9)
	t.Log(strStr("aabaaabaaac", "aabaaac"), "aabaaac", 4)
	t.Log(strStr("adcadcaddcadde", "adcadde"), "adcadde", -1)

}

func TestIndexOfKMPSearch(t *testing.T) {
	t.Log(SubstringIndexsOf("aabaaabaaac", "aabaaac"), "aabaaac", 1)
	t.Log(SubstringIndexsOf("adcadcaddcadde", "adcadde"), "adcadde", 0)
	t.Log(SubstringIndexsOf("ababccbabab", "abab"), "abab", 2)
	t.Log(SubstringIndexsOf("ababccbabab", "ab"), "ab", 4)
	t.Log(SubstringIndexsOf("ababccbabab", "b"), "b", 5)
}

func TestMyAtoi(t *testing.T) {
	t.Log(myAtoi("42"), "42", 42)
	t.Log(myAtoi("9223372036854775808"), "9223372036854775808", 2147483647)
	t.Log(myAtoi("-010"), "-010", -10)
	t.Log(myAtoi("01"), "01", 1)
	t.Log(myAtoi("+1"), "+1", 1)
}

func TestIsInterleave(t *testing.T) {
	t.Log(isInterleave("aabcc", "dbbca", "aadbbcbcac"), true)
	// t.Log(isInterleave("aabcc", "dbbca", "aadbbbaccc"), false)
	// t.Log(isInterleave("", "", ""), true)
}
