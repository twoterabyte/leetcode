/*
 * @lc app=leetcode.cn id=19 lang=golang
 *
 * [19] 删除链表的倒数第 N 个结点
 *
 * https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/description/
 *
 * algorithms
 * Medium (42.59%)
 * Likes:    1538
 * Dislikes: 0
 * Total Accepted:    490.2K
 * Total Submissions: 1.1M
 * Testcase Example:  '[1,2,3,4,5]\n2'
 *
 * 给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
 *
 * 进阶：你能尝试使用一趟扫描实现吗？
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：head = [1,2,3,4,5], n = 2
 * 输出：[1,2,3,5]
 *
 *
 * 示例 2：
 *
 *
 * 输入：head = [1], n = 1
 * 输出：[]
 *
 *
 * 示例 3：
 *
 *
 * 输入：head = [1,2], n = 1
 * 输出：[1]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 链表中结点的数目为 sz
 * 1
 * 0
 * 1
 *
 *
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 双指针解决
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 快指针
	fast := head
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	// 如果快指针已经指向 nil，说明删除的是第一个
	if fast == nil {
		return head.Next
	}
	// 慢指针的前驱
	slowPre := head
	fast = fast.Next
	for fast != nil {
		fast = fast.Next
		slowPre = slowPre.Next
	}
	slowPre.Next = slowPre.Next.Next
	return head
}

// @lc code=end

