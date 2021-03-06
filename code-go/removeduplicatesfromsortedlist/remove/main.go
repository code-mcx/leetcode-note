package remove

import (
    "fmt"
    "strings"
)

// 删除排序链表中的重复元素
// 给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。

// 示例
// 输入: 1->1->2
// 输出: 1->2

// 输入: 1->1->2->3->3
// 输出: 1->2->3

type ListNode struct {
    Val int
    Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
    current, prevNode := head, head
    for current != nil {
        // 找到下一个不相同的节点
        for current.Next != nil && current.Val == current.Next.Val {
            current = current.Next
        }

        // 指向下一个不相同的节点
        prevNode.Next = current.Next

        // 记录上一个节点
        prevNode = current.Next
        // 从下一个不相同的节点继续走
        current = current.Next
    }
    return head
}

// 打印链表
func printListNode(node *ListNode) {
    var str = ""
    for node != nil {
        str += fmt.Sprintf("%d->", node.Val)
        node = node.Next
    }
    str = strings.TrimSuffix(str, "->")
    fmt.Println(str)
}
