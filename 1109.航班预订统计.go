/*
 * @lc app=leetcode.cn id=1109 lang=golang
 *
 * [1109] 航班预订统计
 *
 * https://leetcode-cn.com/problems/corporate-flight-bookings/description/
 *
 * algorithms
 * Medium (50.39%)
 * Likes:    168
 * Dislikes: 0
 * Total Accepted:    29.6K
 * Total Submissions: 58.1K
 * Testcase Example:  '[[1,2,10],[2,3,20],[2,5,25]]\n5'
 *
 * 这里有 n 个航班，它们分别从 1 到 n 进行编号。
 *
 * 有一份航班预订表 bookings ，表中第 i 条预订记录 bookings[i] = [firsti, lasti, seatsi] 意味着在从
 * firsti 到 lasti （包含 firsti 和 lasti ）的 每个航班 上预订了 seatsi 个座位。
 *
 * 请你返回一个长度为 n 的数组 answer，其中 answer[i] 是航班 i 上预订的座位总数。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：bookings = [[1,2,10],[2,3,20],[2,5,25]], n = 5
 * 输出：[10,55,45,25,25]
 * 解释：
 * 航班编号        1   2   3   4   5
 * 预订记录 1 ：   10  10
 * 预订记录 2 ：       20  20
 * 预订记录 3 ：       25  25  25  25
 * 总座位数：      10  55  45  25  25
 * 因此，answer = [10,55,45,25,25]
 *
 *
 * 示例 2：
 *
 *
 * 输入：bookings = [[1,2,10],[2,2,15]], n = 2
 * 输出：[10,25]
 * 解释：
 * 航班编号        1   2
 * 预订记录 1 ：   10  10
 * 预订记录 2 ：       15
 * 总座位数：      10  25
 * 因此，answer = [10,25]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * 1
 * bookings[i].length == 3
 * 1 i i
 * 1 i
 *
 *
 */

// @lc code=start
// 差分数组 sub[i] 代表数组元素 arr[i]-arr[i-1]
// 如果知道 arr[-1] 和 sub 就可以计算出 arr
// 考虑差分数组的含义，假如差分数组只有 sub[i] = x，那么表示从 arr[i:] 中的元素比 arr[:i] 大 x
// 假如只有 sub[i] = x，sub[j]=-x，j>i，表示 arr[i:j] 比 arr 中其他元素是大 x
// 那么本题 [first, last, seats]，表示的差分数组 sub[first-1]=seats, sub[last]=-seats
// 差分数组可以叠加，所以等于对差分数组求和后，再算出 answer
func corpFlightBookings(bookings [][]int, n int) []int {
	// 多 1，就不用判断 last 是否越界
	sub := make([]int, n+1)
	for _, booking := range bookings {
		sub[booking[0]-1] += booking[2]
		sub[booking[1]] -= booking[2]
	}
	ans := make([]int, n)
	sum := 0
	for i := 0; i < n; i++ {
		sum += sub[i]
		ans[i] = sum
	}
	return ans
}

// func corpFlightBookings(bookings [][]int, n int) []int {
// 	ans := make([]int, n)
// 	for _, booking := range bookings {
// 		for i := booking[0] - 1; i < booking[1]; i++ {
// 			ans[i] += booking[2]
// 		}
// 	}
// 	return ans
// }

// @lc code=end

