/*
 * @lc app=leetcode.cn id=457 lang=golang
 *
 * [457] 环形数组是否存在循环
 *
 * https://leetcode-cn.com/problems/circular-array-loop/description/
 *
 * algorithms
 * Medium (43.36%)
 * Likes:    160
 * Dislikes: 0
 * Total Accepted:    27.9K
 * Total Submissions: 64K
 * Testcase Example:  '[2,-1,1,2,2]'
 *
 * 存在一个不含 0 的 环形 数组 nums ，每个 nums[i] 都表示位于下标 i 的角色应该向前或向后移动的下标个数：
 *
 *
 * 如果 nums[i] 是正数，向前（下标递增方向）移动 |nums[i]| 步
 * 如果 nums[i] 是负数，向后（下标递减方向）移动 |nums[i]| 步
 *
 *
 * 因为数组是 环形 的，所以可以假设从最后一个元素向前移动一步会到达第一个元素，而第一个元素向后移动一步会到达最后一个元素。
 *
 * 数组中的 循环 由长度为 k 的下标序列 seq 标识：
 *
 *
 * 遵循上述移动规则将导致一组重复下标序列 seq[0] -> seq[1] -> ... -> seq[k - 1] -> seq[0] ->
 * ...
 * 所有 nums[seq[j]] 应当不是 全正 就是 全负
 * k > 1
 *
 *
 * 如果 nums 中存在循环，返回 true ；否则，返回 false 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [2,-1,1,2,2]
 * 输出：true
 * 解释：存在循环，按下标 0 -> 2 -> 3 -> 0 。循环长度为 3 。
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [-1,2]
 * 输出：false
 * 解释：按下标 1 -> 1 -> 1 ... 的运动无法构成循环，因为循环的长度为 1 。根据定义，循环的长度必须大于 1 。
 *
 *
 * 示例 3:
 *
 *
 * 输入：nums = [-2,1,-1,-2,-2]
 * 输出：false
 * 解释：按下标 1 -> 2 -> 1 -> ... 的运动无法构成循环，因为 nums[1] 是正数，而 nums[2] 是负数。
 * 所有 nums[seq[j]] 应当不是全正就是全负。
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 5000
 * -1000 <= nums[i] <= 1000
 * nums[i] != 0
 *
 *
 *
 *
 * 进阶：你能设计一个时间复杂度为 O(n) 且额外空间复杂度为 O(1) 的算法吗？
 *
 */

// @lc code=start
// 说人的的题目描述：
// 有数组 nums，nums[i] 表示位于下标 i 的角色应该向前或向后移动的下标个数
// 求是否存在角色的子序列使得，在 nums 表示的移动下，是循环移动的。
// next = ((cur + nums[cur])%len(nums) + len(nums)) % len(nums)

// 那么就有遍历每个下标 i 的角色，用快慢指针判断是否成环，且环大于 1
// 时间复杂度是 O(n^2)

// 快慢指针可以判断路径上是否存在还，假如遍历 i 发现没有形成环，那么当遍历 i + 1 时
// 如果进入到 i 的路径，则一定不能形成环
// 所以可以增加一个数组表示已经遍历过的点，再遍历的时候可以直接排除
// 可以通过快指针判断是否进入之前的路径（为什么是快指针，因为慢指针一定是走快指针的路）
// 通过慢指针走过的路径标记已经遍历过
func circularArrayLoop(nums []int) bool {
	n := len(nums)
	// 下一个索引
	next := func(cur int) int {
		return ((cur+nums[cur])%n + n) % n
	}
	for i, num := range nums {
		// 看以当前 i 为起点是否可以出现环
		if num == 0 {
			// 在之前的遍历已经遍历过，且没有出现环（因为如果出现环就返回 true 了）
			continue
		}
		// 快慢指针
		slow := i
		fast := i
		// 如果已经遍历过，则会出现0，导致乘出来的结果是 0，不满足判断
		// 如果方向改变，则乘出来的结果是负数，不满足判断
		// 第一次遍历没有 0，为什么一定跳出循环呢？
		//   因为必定是方向改变了，如果方向不改变，循环最多遍历 n+1 次就会遇到相同的
		//   此时如果方向不改变，则说明出现循环了，所有只有方向改变
		for num*nums[next(fast)] > 0 && num*nums[next(next(fast))] > 0 {
			// 快指针移动第两次
			fast = next(next(fast))
			// 慢指针移动一次
			slow = next(slow)
			// 快指针和慢指针不是指向同一个，继续移动
			if fast != slow {
				continue
			}
			// 快慢指针相遇，说明遇见了环
			if slow == next(slow) {
				// 但是环只有一个元素，不满足
				break
			}
			// 形成环，且不止一个元素，满足要求
			return true
		}
		// slow 走过的部分标记走过
		for j := next(i); num*nums[j] > 0; j = next(j) {
			nums[j] = 0
		}
		nums[i] = 0
	}
	return false
}

// @lc code=end

