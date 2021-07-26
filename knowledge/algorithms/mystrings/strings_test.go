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
