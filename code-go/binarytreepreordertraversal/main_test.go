package binarytreepreordertraversal

import "testing"

func TestPreorderTraversal(t *testing.T) {
    node1 := &TreeNode{Val: 1}
    node2 := &TreeNode{Val: 2}
    node3 := &TreeNode{Val: 3}
    node1.Right = node2
    node2.Left = node3

    results := preorderTraversal(node1)

    t.Log(results)
}
