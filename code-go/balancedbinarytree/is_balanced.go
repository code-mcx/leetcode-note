package main

import "fmt"

// 平衡二叉树
// 给定一个二叉树，判断它是否是高度平衡的二叉树。
// 一棵高度平衡二叉树定义为：一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。

// 示例
// 输入: [3,9,20,null,null,15,7]
//     3
//    / \
//   9  20
//     /  \
//    15   7
// 输出: true

// 输入: [1,2,2,3,3,null,null,4,4]
//        1
//       / \
//      2   2
//     / \
//    3   3
//   / \
//  4   4
// 输出: false

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
    if root == nil {
        return true
    }
    // 判断当前节点的左右子树的高度差
    if abs(height(root.Left) - height(root.Right)) > 1 {
        return false
    }
    // 判断左子节点是否平衡
    isTrue := isBalanced(root.Left)
    if !isTrue {
        return false
    }
    // 判断右子节点是否平衡
    return isBalanced(root.Right)
}

func height(node *TreeNode) int {
    if node == nil {
        return 0
    }
    // 左子树高度
    lHeight := height(node.Left)
    // 右子树高度
    rHeight := height(node.Right)
    // 返回最大高度, 当前节点的高度为 1
    return max(lHeight, rHeight) + 1
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func abs(x int) int {
    if x < 0 {
        return -1 * x
    }
    return x
}

func infixOrder(node *TreeNode) {
    if node == nil {
        return
    }
    infixOrder(node.Left)
    fmt.Println(node.Val)
    infixOrder(node.Right)
}

func main() {
    node1 := &TreeNode{Val: 3}
    node2 := &TreeNode{Val: 9}
    node3 := &TreeNode{Val: 20}
    node4 := &TreeNode{Val: 15}
    node5 := &TreeNode{Val: 7}
    node1.Left = node2
    node1.Right = node3
    node3.Left = node4
    node3.Right = node5

    infixOrder(node1)

    result := isBalanced(node1)

    fmt.Println(result)

    node1 = &TreeNode{Val: 1}
    node2 = &TreeNode{Val: 2}
    node3 = &TreeNode{Val: 2}
    node4 = &TreeNode{Val: 3}
    node5 = &TreeNode{Val: 3}
    node6 := &TreeNode{Val: 4}
    node7 := &TreeNode{Val: 4}
    node1.Left = node2
    node1.Right = node3
    node2.Left = node4
    node2.Right = node5
    node4.Left = node6
    node4.Right = node7

    infixOrder(node1)

    result = isBalanced(node1)

    fmt.Println(result)
}
