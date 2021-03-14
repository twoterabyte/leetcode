/*
 * @lc app=leetcode.cn id=111 lang=cpp
 *
 * [111] 二叉树的最小深度
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 * };
 */
#include <queue>
using namespace std;

class Solution {
public:
    int minDepth(TreeNode* root) {
        // 无返回 0
        if (!root) {
            return 0;
        }
        // 初始深度
        int depth = 0;
        // BFS 的队列
        std::queue<TreeNode*> q;
        q.push(root);
        while (!q.empty()) {
            // 深度加 1
            ++depth;
            // 取出一层
            size_t len = q.size();
            for (int i = 0; i < len; ++i) {
                auto *node = q.front();
                // 达到叶子节点
                if (!node->left && !node->right) return depth;
                if (node->left) q.push(node->left);
                if (node->right) q.push(node->right);
                q.pop();
            }
        }
        return depth;
    }
};
// @lc code=end

