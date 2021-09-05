/*
 * @lc app=leetcode.cn id=470 lang=golang
 *
 * [470] 用 Rand7() 实现 Rand10()
 *
 * https://leetcode-cn.com/problems/implement-rand10-using-rand7/description/
 *
 * algorithms
 * Medium (55.30%)
 * Likes:    284
 * Dislikes: 0
 * Total Accepted:    49.7K
 * Total Submissions: 90.9K
 * Testcase Example:  '1'
 *
 * 已有方法 rand7 可生成 1 到 7 范围内的均匀随机整数，试写一个方法 rand10 生成 1 到 10 范围内的均匀随机整数。
 *
 * 不要使用系统的 Math.random() 方法。
 *
 *
 *
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: 1
 * 输出: [7]
 *
 *
 * 示例 2:
 *
 *
 * 输入: 2
 * 输出: [8,4]
 *
 *
 * 示例 3:
 *
 *
 * 输入: 3
 * 输出: [8,1,10]
 *
 *
 *
 *
 * 提示:
 *
 *
 * rand7 已定义。
 * 传入参数: n 表示 rand10 的调用次数。
 *
 *
 *
 *
 * 进阶:
 *
 *
 * rand7()调用次数的 期望值 是多少 ?
 * 你能否尽量少调用 rand7() ?
 *
 *
 */

// 7 进制加拒绝采样
// 生成两次数字范围在 [0,48]，其中 [0,39] 进去 %10+1，则可以得到 rand10，
// 如果 [40,48]，则第二次重新生成，直到在 [0,39] 为止。
// 两次看作一次，40/49概率能获取 rand10，期望次数就是 1/(40/49)，两次看作看一次
// 所以总次数是 2*1/(40/49)=2.45
// 利用被拒绝的 [40,48]，看作 [0,8]，再生成一次，则范围是 [0,62]，
// 拒绝 [60,62]，看作 [0,2]，再生成一次，则范围是[0,20]
// 拒绝 [20,20], 看作 [0,0]，没有作用了，只能从头开始

// @lc code=start
// func rand10() int {
// 	for {
// 		n := (rand7()-1)*7 + rand7() - 1
// 		if n <= 40 {
// 			return n%10 + 1
// 		}
// 	}
// }

func rand10() int {
	for {
		n := (rand7()-1)*7 + rand7() - 1
		if n < 40 {
			return n%10 + 1
		}
		n = n%10*7 + rand7() - 1
		if n < 60 {
			return n%10 + 1
		}
		n = n%10*7 + rand10() - 1
		if n < 20 {
			return n%10 + 1
		}
	}
}

// @lc code=end

