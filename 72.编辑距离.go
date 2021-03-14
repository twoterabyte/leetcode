/*
 * @lc app=leetcode.cn id=72 lang=golang
 *
 * [72] 编辑距离
 *
 * https://leetcode-cn.com/problems/edit-distance/description/
 *
 * algorithms
 * Hard (59.57%)
 * Likes:    1016
 * Dislikes: 0
 * Total Accepted:    72.7K
 * Total Submissions: 121.9K
 * Testcase Example:  '"horse"\n"ros"'
 *
 * 给你两个单词 word1 和 word2，请你计算出将 word1 转换成 word2 所使用的最少操作数 。
 * 
 * 你可以对一个单词进行如下三种操作：
 * 
 * 
 * 插入一个字符
 * 删除一个字符
 * 替换一个字符
 * 
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：word1 = "horse", word2 = "ros"
 * 输出：3
 * 解释：
 * horse -> rorse (将 'h' 替换为 'r')
 * rorse -> rose (删除 'r')
 * rose -> ros (删除 'e')
 * 
 * 
 * 示例 2：
 * 
 * 输入：word1 = "intention", word2 = "execution"
 * 输出：5
 * 解释：
 * intention -> inention (删除 't')
 * inention -> enention (将 'i' 替换为 'e')
 * enention -> exention (将 'n' 替换为 'x')
 * exention -> exection (将 'n' 替换为 'c')
 * exection -> execution (插入 'u')
 * 
 * 
 */

// dp[i][j] 表示将 word1[0:i] 转换成 word2[0:j] 所需要的最少操作数
// dp[0][0]=0，dp[0][j]=j，dp[i][0]=i
// dp[i][j] = 
// if word1[i-1] == word2[j-1] then dp[i-1][j-1]
// if min{dp[i-1][j-1]+1, dp[i-1][j]+1, dp[i][j-1]}
// dp[i-1][j-1]+1 表示替换一个字符，那么只需要看 word1[0:i-1] 和 word2[0:j-1]
// dp[i-1][j]+1 表示删除一个字符，那么只需要看 word1[0:i-1] 和 word2[0:j]
// dp[i][j-1]+1 表示插入一个字符word2[j-1]，那么只需要看 word1[0:i] 和 word2[0:j-1]
// @lc code=start
func minDistance(word1 string, word2 string) int {
	// dp 用状态压缩，只需要两个数组交替
	dp := make([]int, len(word2)+1)
	dpPre := make([]int, len(word2)+1)
	// dp[0][j]=j 初始化
	for j := range dp {
		dp[j] = j
	}
	for i := 1; i <= len(word1); i++ {
		// 交换数组
		dp, dpPre = dpPre, dp
		// dp[i][0]=i 初始化
		dp[0] = i
		for j := 1; j <= len(word2); j++ {
			if (word1[i-1] == word2[j-1]) {
				dp[j] = dpPre[j-1]
				continue
			}
			dp[j] = min3(dpPre[j-1]+1, dpPre[j]+1, dp[j-1]+1)
		}
	}
	return dp[len(word2)]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func min3(x, y, z int) int {
	return min(min(x, y), z)
}
// @lc code=end

