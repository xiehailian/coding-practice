package basic

import (
	"fmt"
	"go/algo/basic"
)

// 面试题6：从尾到头打印链表
// 题目：输入一个链表的头结点，从尾到头反过来打印出每个结点的值。


func PrintListInReversedOrder(head *basic.Element) {
	if head != nil {
		PrintListInReversedOrder(head.Next())
		fmt.Printf("%d -> ", head.Value)
	}
}