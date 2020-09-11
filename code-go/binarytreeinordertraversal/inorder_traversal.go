package main

import "fmt"

// 二叉树的中序遍历
// 给定一个二叉树，返回它的中序遍历。

// 示例
// 输入: [1,null,2,3]
//   1
//    \
//     2
//    /
//   3
// 输出: [1,3,2]

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
    // 非递归
    return inorderNoRecursion(root)
    // 递归
    // return inorderRecursion(root)
}

func inorderNoRecursion(node *TreeNode) []int {
    var stack []*TreeNode
    output := make([]int, 0)
    if node == nil {
        return output
    }
    // 将当前节点压入栈
    stack = append(stack, node)
    // 记录每次迭代时的节点
    current := node
    for len(stack) > 0 {
        // 将当前节点的所有最左边的节点入栈
        for current.Left != nil {
            // 将最左边的节点压入栈
            stack = append(stack, current.Left)
            current = current.Left
        }

        // 弹出当前节点
        current = stack[len(stack) - 1]
        stack = stack[:len(stack) - 1]

        output = append(output, current.Val)

        // 当前节点的右子节点不为空, 重复循环
        if current.Right != nil {
            // 将右子节点压入栈
            stack = append(stack, current.Right)
            current = current.Right
        } else {
            // 当前节点的右子节点为空, 赋值为一个新的节点, 避免继续重复将最左边的节点压入栈
            current = &TreeNode{Val: -1}
        }
    }

    return output
}

func inorderRecursion(root *TreeNode) []int {
    if root == nil {
        return []int{}
    }
    nodes := make([]int, 0)
    // 先访问左子节点
    nodes = append(nodes, inorderRecursion(root.Left)...)
    // 再将当前节点存入数组
    nodes = append(nodes, root.Val)
    // 最后访问右子节点
    nodes = append(nodes, inorderRecursion(root.Right)...)
    return nodes
}

func main() {
    node1 := &TreeNode{Val: 1}
    node2 := &TreeNode{Val: 2}
    node3 := &TreeNode{Val: 3}
    node1.Right = node2
    node2.Left = node3

    results := inorderTraversal(node1)

    fmt.Println(results)
}
