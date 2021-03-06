package com.dxx.palindromenum;

/**
 * 回文数
 * 判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
 *
 * 示例
 * 输入: 121
 * 输出: true
 *
 * 输入: -121
 * 输出: false
 * 解释: 从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
 *
 * 输入: 10
 * 输出: false
 * 解释: 从右向左读, 为 01 。因此它不是一个回文数。
 */
public class Main {

    public static boolean isPalindrome(int num) {
        if (num < 0) {
            return false;
        }

        // 以下代码反转整数，不考虑越界
        int revNum = 0;
        int digit = 0;
        int n = num;
        while (n != 0) {
            digit = n % 10;
            n /= 10;
            revNum = revNum * 10 + digit;
        }

        return num == revNum;
    }

    public static void main(String[] args) {
        int num = 121;
        boolean result = isPalindrome(num);
        System.out.println(result);

        num = -121;
        result = isPalindrome(num);
        System.out.println(result);

        num = 10;
        result = isPalindrome(num);
        System.out.println(result);
    }
}
