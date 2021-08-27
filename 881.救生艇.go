/*
 * @lc app=leetcode.cn id=881 lang=golang
 *
 * [881] 救生艇
 *
 * https://leetcode-cn.com/problems/boats-to-save-people/description/
 *
 * algorithms
 * Medium (47.72%)
 * Likes:    175
 * Dislikes: 0
 * Total Accepted:    44.3K
 * Total Submissions: 82.2K
 * Testcase Example:  '[1,2]\n3'
 *
 * 第 i 个人的体重为 people[i]，每艘船可以承载的最大重量为 limit。
 *
 * 每艘船最多可同时载两人，但条件是这些人的重量之和最多为 limit。
 *
 * 返回载到每一个人所需的最小船数。(保证每个人都能被船载)。
 *
 *
 *
 * 示例 1：
 *
 * 输入：people = [1,2], limit = 3
 * 输出：1
 * 解释：1 艘船载 (1, 2)
 *
 *
 * 示例 2：
 *
 * 输入：people = [3,2,2,1], limit = 3
 * 输出：3
 * 解释：3 艘船分别载 (1, 2), (2) 和 (3)
 *
 *
 * 示例 3：
 *
 * 输入：people = [3,5,3,4], limit = 5
 * 输出：4
 * 解释：4 艘船分别载 (3), (3), (4), (5)
 *
 * 提示：
 *
 *
 * 1 <= people.length <= 50000
 * 1 <= people[i] <= limit <= 30000
 *
 *
 */

// @lc code=start
// 先对 people 进行排序
// 1. 取最轻的人，看是否能和最重的人组合
//   * 如果不能，说明重的人不可能和任何组合，需要自己一艘船
//   * 如果能，则和最重的人组合
// 2. 继续第一步
func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)
	i := 0               // 最轻的人
	j := len(people) - 1 // 最重的人
	cnt := 0
	// 最后如果只有一个人，虽然最轻和最重都是同一个人，但是 cnt++，答案是正确的
	for i <= j {
		// 无论是否组合都需要一首艘船
		cnt++
		if people[i]+people[j] > limit {
			// 明重的人不可能和任何组合，需要自己一艘船
			j--
			continue
		}
		// 两人组成一艘船
		i++
		j--
	}
	return cnt
}

// @lc code=end

