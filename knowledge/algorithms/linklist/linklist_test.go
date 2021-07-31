//  linklist_test.go
//
//
//  Created by JohnnyB0Y on 2021/07/31.
//  Copyright © 2021 JohnnyB0Y. All rights reserved.

package linklist

import "testing"

func TestMergeTwoLists(t *testing.T) {
	l1 := ListNode{1, nil}
	l1.Next = &ListNode{2, nil}

	l2 := ListNode{30, nil}
	l2.Next = &ListNode{40, nil}
	links := mergeTwoLists(&l1, &l2)
	for links != nil {
		t.Log(links.Val)
		links = links.Next
	}
}
