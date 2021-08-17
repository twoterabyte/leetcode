/*
 * @lc app=leetcode.cn id=526 lang=golang
 *
 * [526] 优美的排列
 *
 * https://leetcode-cn.com/problems/beautiful-arrangement/description/
 *
 * algorithms
 * Medium (65.64%)
 * Likes:    147
 * Dislikes: 0
 * Total Accepted:    12.7K
 * Total Submissions: 19.2K
 * Testcase Example:  '2'
 *
 * 假设有从 1 到 N 的 N 个整数，如果从这 N 个数字中成功构造出一个数组，使得数组的第 i 位 (1 <= i <= N)
 * 满足如下两个条件中的一个，我们就称这个数组为一个优美的排列。条件：
 *
 *
 * 第 i 位的数字能被 i 整除
 * i 能被第 i 位上的数字整除
 *
 *
 * 现在给定一个整数 N，请问可以构造多少个优美的排列？
 *
 * 示例1:
 *
 *
 * 输入: 2
 * 输出: 2
 * 解释:
 *
 * 第 1 个优美的排列是 [1, 2]:
 * ⁠ 第 1 个位置（i=1）上的数字是1，1能被 i（i=1）整除
 * ⁠ 第 2 个位置（i=2）上的数字是2，2能被 i（i=2）整除
 *
 * 第 2 个优美的排列是 [2, 1]:
 * ⁠ 第 1 个位置（i=1）上的数字是2，2能被 i（i=1）整除
 * ⁠ 第 2 个位置（i=2）上的数字是1，i（i=2）能被 1 整除
 *
 *
 * 说明:
 *
 *
 * N 是一个正整数，并且不会超过15。
 *
 *
 */

// 回溯算法：
// 枚举每个位置上可能值，直到全部枚举完
// 时间复杂度 O(n!)

// 可以优化的点，每个位置 i 上不需要枚1-N，只需要 i*k <= 15，k是正整数 或者 i/k > 0 && i%k == 0，当 k 是正整数，且
// @lc code=start
func countArrangement(n int) int {
	useds := make([]bool, n)
	// 至少有 1，即空排列
	cnt := 0
	var backtrack func(index int)
	// 在 array[:index] 已有排列的情况下，确定 index 的值
	backtrack = func(index int) {
		// 递归终止条件
		if index == n {
			cnt++
			return
		}
		for i, used := range useds {
			if used {
				continue
			}
			// 当前数不满足要求，就直接跳过了
			if !valid(index+1, i+1) {
				continue
			}
			// num 可以加入到当前排列下
			useds[i] = true
			backtrack(index + 1)
			useds[i] = false
		}
	}
	backtrack(0)
	return cnt
}

func valid(index, num int) bool {
	if num%index == 0 {
		return true
	}
	if index%num == 0 {
		return true
	}
	return false
}

// @lc code=end

