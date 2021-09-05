/*
 * @lc app=leetcode.cn id=215 lang=golang
 *
 * [215] 数组中的第K个最大元素
 *
 * https://leetcode-cn.com/problems/kth-largest-element-in-an-array/description/
 *
 * algorithms
 * Medium (64.72%)
 * Likes:    1258
 * Dislikes: 0
 * Total Accepted:    421.9K
 * Total Submissions: 652.2K
 * Testcase Example:  '[3,2,1,5,6,4]\n2'
 *
 * 给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。
 *
 * 请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: [3,2,1,5,6,4] 和 k = 2
 * 输出: 5
 *
 *
 * 示例 2:
 *
 *
 * 输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
 * 输出: 4
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * -10^4
 *
 *
 */

// @lc code=start

// 用快排的分割函数，将 nums 分割为两部分，如果刚好分割的点等于 k，则是要求的值
import "math/rand"

func findKthLargest(nums []int, k int) int {
	// 随机取一个数字作为分割
	num := rand.Int() % len(nums)
	// 将分割元素交换到最后
	nums[len(nums)-1], nums[num] = nums[num], nums[len(nums)-1]
	pivot := nums[len(nums)-1]
	// 分割后，nums[p-1] 是 nums[pivot]，nums[:p-1] 大于 nums[pivot]，nums[p:] 小于 nums[pivot]
	p := partion(nums, func(i int) bool { return nums[i] >= pivot })
	if p == k {
		return nums[p-1]
	} else if p < k {
		return findKthLargest(nums[p:], k-p)
	} else {
		return findKthLargest(nums[:p-1], k)
	}
}

func partion(slices []int, fn func(i int) bool) int {
	// slices[:i] 都是满足 fn 为 true
	i := 0
	for j := 0; j < len(slices); j++ {
		// 如果当前值满足换到 i 位置，同时 i 递增
		if fn(j) {
			slices[i], slices[j] = slices[j], slices[i]
			i++
		}
	}
	return i
}

// @lc code=end

