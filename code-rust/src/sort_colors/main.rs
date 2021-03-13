/// 颜色分类

/// 给定一个包含红色、白色和蓝色，一共 n 个元素的数组，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。
/// 此题中，我们使用整数 0、1 和 2 分别表示红色、白色和蓝色。

/// 注意: 不能使用代码库中的排序函数来解决这道题。

/// 进阶:
/// 一个直观的解决方案是使用计数排序的两趟扫描算法。
/// 首先，迭代计算出0、1 和 2 元素的个数，然后按照0、1、2的排序，重写当前数组。
/// 你能想出一个仅使用常数空间的一趟扫描算法吗？

/// 示例
/// 输入: [2, 0, 2, 1, 1, 0]
/// 输出: [0, 0, 1, 1, 2, 2]

pub struct Solution {}

impl Solution {
    pub fn sort_colors(nums: &mut Vec<i32>) {
        // p0 表示 0 要放置的下标, p1 表示 1 要放置的下标
        let (mut p0, mut p1) = (0, 0);
        for i in 0..nums.len() {
            if nums[i] == 0 {
                // 交换 i 和 p0 指向的元素
                Solution::swap(nums, i, p0);
                // 有可能 p0 指向的元素为 1, 跟 i 发生了交换
                if p0 < p1 {
                    // 将 p1 和 i 交换
                    Solution::swap(nums, i, p1);
                }
                p0 += 1;
                p1 += 1;
            } else if nums[i] == 1 {
                // 交换 i 和 p1 指向的元素
                Solution::swap(nums, i, p1);
                p1 += 1;
            }
        }
    }

    fn swap(nums: &mut Vec<i32>, a: usize, b: usize) {
        let temp = nums[a];
        nums[a] = nums[b];
        nums[b] = temp;
    }
}

#[test]
fn test_sort_colors() {
    let mut nums = vec![2, 0, 2, 1, 1, 0];
    println!("{:?}", nums);

    Solution::sort_colors(&mut nums);
    println!("{:?}", nums);
}
