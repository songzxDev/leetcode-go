package week01

import (
	"sort"
)

//给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
//
// 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。
//
//
//
// 示例:
//
// 给定 nums = [2, 7, 11, 15], target = 9
//
//因为 nums[0] + nums[1] = 2 + 7 = 9
//所以返回 [0, 1]
//
// Related Topics 数组 哈希表

//leetcode submit region begin(Prohibit modification and deletion)
func twoSum(nums []int, target int) []int {
	myCache := make(map[int]int)
	for i, num := range nums {
		if v, ok := myCache[target-num]; ok {
			return []int{v, i}
		}
		myCache[num] = i
	}
	return []int{}
}

//leetcode submit region end(Prohibit modification and deletion)

//给定两个单词（beginWord 和 endWord）和一个字典，找到从 beginWord 到 endWord 的最短转换序列的长度。转换需遵循如下规则：
//
//
//
// 每次转换只能改变一个字母。
// 转换过程中的中间单词必须是字典中的单词。
//
//
// 说明:
//
//
// 如果不存在这样的转换序列，返回 0。
// 所有单词具有相同的长度。
// 所有单词只由小写字母组成。
// 字典中不存在重复的单词。
// 你可以假设 beginWord 和 endWord 是非空的，且二者不相同。
//
//
// 示例 1:
//
// 输入:
//beginWord = "hit",
//endWord = "cog",
//wordList = ["hot","dot","dog","lot","log","cog"]
//
//输出: 5
//
//解释: 一个最短转换序列是 "hit" -> "hot" -> "dot" -> "dog" -> "cog",
//     返回它的长度 5。
//
//
// 示例 2:
//
// 输入:
//beginWord = "hit"
//endWord = "cog"
//wordList = ["hot","dot","dog","lot","log"]
//
//输出: 0
//
//解释: endWord "cog" 不在字典中，所以无法进行转换。
// Related Topics 广度优先搜索

//leetcode submit region begin(Prohibit modification and deletion)
func ladderLength(beginWord string, endWord string, wordList []string) int {
	if wordList != nil && len(wordList) > 0 {
		wordSet, beginSet, endSet, steps := getWordMap(wordList), getWordMap([]string{beginWord}), getWordMap([]string{endWord}), 1
		if _, ok := wordSet[endWord]; !ok {
			return 0
		}
		for len(beginSet) > 0 {
			steps++
			nextSet := make(map[string]bool)
			if len(beginSet) > len(endSet) {
				beginSet, endSet = endSet, beginSet
			}
			for word := range beginSet {
				for i := 0; i < len(word); i++ {
					for _, c := range "abcdefghijklmnopqrstuvwxyz" {
						if c != rune(word[i]) {
							target := word[:i] + string(c) + word[i+1:]
							if _, ok := endSet[target]; ok {
								return steps
							}
							if _, ok := wordSet[target]; ok {
								delete(wordSet, target)
								nextSet[target] = true
							}
						}
					}
				}
			}
			beginSet = nextSet
		}
	}
	return 0
}

func getWordMap(wordList []string) map[string]bool {
	wordSet := make(map[string]bool)
	if wordList != nil && len(wordList) > 0 {
		for _, word := range wordList {
			wordSet[word] = true
		}
	}
	return wordSet
}

//leetcode submit region end(Prohibit modification and deletion)

//给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复
//的三元组。
//
// 注意：答案中不可以包含重复的三元组。
//
//
//
// 示例：
//
// 给定数组 nums = [-1, 0, 1, 2, -1, -4]，
//
//满足要求的三元组集合为：
//[
//  [-1, 0, 1],
//  [-1, -1, 2]
//]
//
// Related Topics 数组 双指针

//leetcode submit region begin(Prohibit modification and deletion)
func threeSum(nums []int) [][]int {
	var res [][]int
	if nums == nil || len(nums) < 3 {
		return res
	}
	sort.Ints(nums)
	n, first := len(nums), nums[0]
	for i := 0; first <= 0 && i < n-2; i++ {
		if i == 0 || nums[i] > nums[i-1] {
			j, k := i+1, n-1
			for j < k {
				afterAdd := nums[i] + nums[j] + nums[k]
				if afterAdd == 0 {
					res = append(res, []int{nums[i], nums[j], nums[k]})
					j++
					k--
					for j < k && nums[j] == nums[j-1] {
						j++
					}
					for j < k && nums[k] == nums[k+1] {
						k--
					}
				} else if afterAdd < 0 {
					j++
				} else {
					k--
				}
			}
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
//数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
//
//
//
// 示例：
//
// 输入：n = 3
//输出：[
//       "((()))",
//       "(()())",
//       "(())()",
//       "()(())",
//       "()()()"
//     ]
//
// Related Topics 字符串 回溯算法

//leetcode submit region begin(Prohibit modification and deletion)
func generateParenthesis(n int) []string {
	var helper func(left int, right int, s string, res *[]string, n int) []string
	helper = func(left int, right int, s string, res *[]string, n int) []string {
		if left+right == (n << 1) {
			*res = append(*res, s)
			return *res
		}
		if left < n {
			helper(left+1, right, s+"(", res, n)
		}
		if right < left {
			helper(left, right+1, s+")", res, n)
		}
		return *res
	}
	return helper(0, 0, "", &[]string{}, n)
}

//leetcode submit region end(Prohibit modification and deletion)

//给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。
//
// 示例:
//
// 输入: ["eat", "tea", "tan", "ate", "nat", "bat"]
//输出:
//[
//  ["ate","eat","tea"],
//  ["nat","tan"],
//  ["bat"]
//]
//
// 说明：
//
//
// 所有输入均为小写字母。
// 不考虑答案输出的顺序。
//
// Related Topics 哈希表 字符串

//leetcode submit region begin(Prohibit modification and deletion)
func groupAnagrams(strs []string) [][]string {
	var res [][]string
	if strs == nil || len(strs) == 0 {
		return res
	}
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101}
	getNumKey := func(stt string) int {
		numKey := 1
		for _, v := range stt {
			numKey *= primes[v-'a']
		}
		return numKey
	}
	countMap := make(map[int][]string)
	for _, stt := range strs {
		numKey := getNumKey(stt)
		countMap[numKey] = append(countMap[numKey], stt)
	}
	for _, value := range countMap {
		res = append(res, value)
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

//给定一个二叉树，返回它的中序 遍历。
//
// 示例:
//
// 输入: [1,null,2,3]
//   1
//    \
//     2
//    /
//   3
//
//输出: [1,3,2]
//
// 进阶: 递归算法很简单，你可以通过迭代算法完成吗？
// Related Topics 栈 树 哈希表

//leetcode submit region begin(Prohibit modification and deletion)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var treeHelper func(node *TreeNode, res *[]int) []int
	treeHelper = func(node *TreeNode, res *[]int) []int {
		if node == nil {
			return *res
		}
		treeHelper(node.Left, res)
		*res = append(*res, node.Val)
		treeHelper(node.Right, res)
		return *res
	}
	return treeHelper(root, &[]int{})
}

//leetcode submit region end(Prohibit modification and deletion)

//给定一个二叉树，返回它的 前序 遍历。
//
// 示例:
//
// 输入: [1,null,2,3]
//   1
//    \
//     2
//    /
//   3
//
//输出: [1,2,3]
//
//
// 进阶: 递归算法很简单，你可以通过迭代算法完成吗？
// Related Topics 栈 树

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
	var treeHelper func(node *TreeNode, res *[]int) []int
	treeHelper = func(node *TreeNode, res *[]int) []int {
		if node == nil {
			return *res
		}
		*res = append(*res, node.Val)
		treeHelper(node.Left, res)
		treeHelper(node.Right, res)
		return *res
	}
	return treeHelper(root, &[]int{})
}

//leetcode submit region end(Prohibit modification and deletion)
