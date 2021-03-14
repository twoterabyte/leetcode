/*
 * @lc app=leetcode.cn id=494 lang=golang
 *
 * [494] 目标和
 *
 * https://leetcode-cn.com/problems/target-sum/description/
 *
 * algorithms
 * Medium (44.40%)
 * Likes:    323
 * Dislikes: 0
 * Total Accepted:    35.9K
 * Total Submissions: 80.9K
 * Testcase Example:  '[1,1,1,1,1]\n3'
 *
 * 给定一个非负整数数组，a1, a2, ..., an, 和一个目标数，S。现在你有两个符号 + 和 -。对于数组中的任意一个整数，你都可以从 + 或
 * -中选择一个符号添加在前面。
 * 
 * 返回可以使最终数组和为目标数 S 的所有添加符号的方法数。
 * 
 * 
 * 
 * 示例：
 * 
 * 输入：nums: [1, 1, 1, 1, 1], S: 3
 * 输出：5
 * 解释：
 * 
 * -1+1+1+1+1 = 3
 * +1-1+1+1+1 = 3
 * +1+1-1+1+1 = 3
 * +1+1+1-1+1 = 3
 * +1+1+1+1-1 = 3
 * 
 * 一共有5种方法让最终目标和为3。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 数组非空，且长度不会超过 20 。
 * 初始的数组的和不会超过 1000 。
 * 保证返回的最终结果能被 32 位整数存下。
 * 
 * 
 */

// Sp - Sn = S
// Sp = S + Sn
// Sp + Sp = S + Sn + Sp
// 2Sp = S + Sall
// Sp = (S + Sall) / 2
// 所以问题转换一个背包问题
// dp[i][j] 表示对于前 i 个数字，当前和为 j, 有 dp[i][j] 种选择方式
// dp[0][:] = 0; dp[0][0] = 1
// dp[i][j] = dp[i-1][j] + dp[i-1][j-nums[i-1]]
// dp[i-1][j] 表示 i 不进行选择
// dp[i-1][j-nums[i-1]] 表示 i 进行选择，nums[i-1] 因为我们 nums 是从 0 开始的
// 目的求dp[len(nums)][Sp]
// @lc code=start

func findTargetSumWays(nums []int, S int) int {
	sAll := 0
	for _, num := range nums {
		sAll += num
	}
	// 显然，目标和不可能超过总和
	if (S > sAll) {
		return 0;
	}
	// 如果不能整除，说明不可能正确
	if (S + sAll) % 2 != 0 {
		return 0
	}
	Sp := (S + sAll) / 2
	// dp 状态压缩，只用一个数组
	dp := make([]int, Sp+1)
	dp[0] = 1
	for i := 1; i <= len(nums); i++ {
		// j 需要从后向前遍历，否则用一维数组会被覆盖
		for j := Sp; j >= 0; j-- {
			if j-nums[i-1] >= 0 {
				dp[j] = dp[j] + dp[j-nums[i-1]]
			}
		}
	}
	return dp[Sp]
}

// 状态压缩后内存还是爆了
// func findTargetSumWays(nums []int, S int) int {
// 	sAll := 0
// 	for _, num := range nums {
// 		sAll += num
// 	}
// 	// 如果不能整除，说明不可能正确
// 	if (S + sAll) % 2 != 0 {
// 		return 0
// 	}
// 	Sp := (S + sAll) / 2
// 	// dp 状态压缩，只用一个数组
// 	dp := make([]int, Sp+1)
// 	dp[0] = 1
// 	for i := 1; i <= len(nums); i++ {
// 		// j 需要从后向前遍历，否则用一维数组会被覆盖
// 		for j := Sp; j > 0; j-- {
// 			if j-nums[i-1] >= 0 {
// 				dp[j] = dp[j] + dp[j-nums[i-1]]
// 			}
// 		}
// 	}
// 	return dp[Sp]
// }

// 内存爆了
// func findTargetSumWays(nums []int, S int) int {
// 	sAll := 0
// 	for _, num := range nums {
// 		sAll += num
// 	}
// 	// 如果不能整除，说明不可能正确
// 	if (S + sAll) % 2 != 0 {
// 		return 0
// 	}
// 	Sp := (S + sAll) / 2
// 	// dp 初始化
// 	dp := make([][]int, len(nums)+1)
// 	for i := range dp {
// 		dp[i] = make([]int, Sp+1)
// 		dp[i][0] = 1
// 	}
// 	for i := 1; i <= len(nums); i++ {
// 		for j := 1; j <= Sp; j++ {
// 			if j-nums[i-1] < 0 {
// 				dp[i][j] = dp[i-1][j]
// 			} else {
// 				dp[i][j] = dp[i-1][j] + dp[i-1][j-nums[i-1]]
// 			}
// 		}
// 	}
// 	return dp[len(nums)][Sp]
// }
// @lc code=end

