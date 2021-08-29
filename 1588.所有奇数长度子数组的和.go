/*
 * @lc app=leetcode.cn id=1588 lang=golang
 *
 * [1588] 所有奇数长度子数组的和
 *
 * https://leetcode-cn.com/problems/sum-of-all-odd-length-subarrays/description/
 *
 * algorithms
 * Easy (80.16%)
 * Likes:    123
 * Dislikes: 0
 * Total Accepted:    39.9K
 * Total Submissions: 47.5K
 * Testcase Example:  '[1,4,2,5,3]'
 *
 * 给你一个正整数数组 arr ，请你计算所有可能的奇数长度子数组的和。
 *
 * 子数组 定义为原数组中的一个连续子序列。
 *
 * 请你返回 arr 中 所有奇数长度子数组的和 。
 *
 *
 *
 * 示例 1：
 *
 * 输入：arr = [1,4,2,5,3]
 * 输出：58
 * 解释：所有奇数长度子数组和它们的和为：
 * [1] = 1
 * [4] = 4
 * [2] = 2
 * [5] = 5
 * [3] = 3
 * [1,4,2] = 7
 * [4,2,5] = 11
 * [2,5,3] = 10
 * [1,4,2,5,3] = 15
 * 我们将所有值求和得到 1 + 4 + 2 + 5 + 3 + 7 + 11 + 10 + 15 = 58
 *
 * 示例 2：
 *
 * 输入：arr = [1,2]
 * 输出：3
 * 解释：总共只有 2 个长度为奇数的子数组，[1] 和 [2]。它们的和为 3 。
 *
 * 示例 3：
 *
 * 输入：arr = [10,11,12]
 * 输出：66
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= arr.length <= 100
 * 1 <= arr[i] <= 1000
 *
 *
 */

// @lc code=start
// dp(i) 表示 arr[:i+1] 的所有奇数长度子数组的和
// dp(0) = arr[0]
// dp(i) = dp(i-1) + sum{sum(arr[i-2k:i+1]) for i-2k>=0 and k >= 0}
// sum(arr[i-2k:i+1]) = sum(arr[:i+1])-sum(arr[:i-2k])
// 时间复杂度：O(n^2)
// 空间复杂度：O(n)
// func sumOddLengthSubarrays(arr []int) int {
// 	// 累计和，用于辅助计算
// 	sum := make([]int, len(arr)+1)
// 	sum[0] := 0
// 	for i, num := range arr {
// 		sum[i] = sum[i-1] + num
// 	}
// 	dp := arr[0]
// 	for i := 2; i < len(arr); i++ {
// 		for j := i; j >= 0; j -= 2 {
// 			dp += sum[i+1] - sum[j]
// 		}
// 	}
// 	return dp
// }

// 考虑 arr[i] 能成为奇数长度子数组，则 arr[i] 左侧和右侧都是奇数，或者都是偶数（0个也是偶数）才可以。
// arr[i] 左侧有 i 个数字，右侧有 len(arr)-1-i 个数字
// 左侧奇数个数 (i+1)/2，偶数个数 i/2+1
// 右侧奇数个数 (len(arr)-i)/2，偶数个数 (len(arr)-1-i)/2+1
// arr[i] 为和贡献了 ((i+1)/2)*((len(arr)-i)/2)+(i/2+1)*((len(arr)-1-i)/2+1) 次

func sumOddLengthSubarrays(arr []int) int {
	ans := 0
	for i, num := range arr {
		l1 := (i + 1) / 2
		l2 := i/2 + 1
		r1 := (len(arr) - i) / 2
		r2 := (len(arr)-1-i)/2 + 1
		ans += (l1*r1 + l2*r2) * num
	}
	return ans
}

// @lc code=end

