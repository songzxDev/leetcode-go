package main

import (
	"fmt"
	"math"
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

//将两个升序链表合并为一个新的升序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
//
// 示例：
//
// 输入：1->2->4, 1->3->4
//输出：1->1->2->3->4->4
//
// Related Topics 链表

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	} else if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	}
}

//leetcode submit region end(Prohibit modification and deletion)

//反转一个单链表。
//
// 示例:
//
// 输入: 1->2->3->4->5->NULL
//输出: 5->4->3->2->1->NULL
//
// 进阶:
//你可以迭代或递归地反转链表。你能否用两种方法解决这道题？
// Related Topics 链表

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	after := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return after
}

//leetcode submit region end(Prohibit modification and deletion)

//给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。
//
// 给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
//
//
//
// 示例:
//
// 输入："23"
//输出：["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
//
//
// 说明:
//尽管上面的答案是按字典序排列的，但是你可以任意选择答案输出的顺序。
// Related Topics 字符串 回溯算法

//leetcode submit region begin(Prohibit modification and deletion)
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	tels := []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	var letterHelper func(digits string, res *[]string, s string, depth int) []string
	letterHelper = func(digits string, res *[]string, s string, depth int) []string {
		if len(s) == len(digits) {
			*res = append(*res, s)
			return *res
		}
		coolie := tels[digits[depth]-'0']
		for _, c := range coolie {
			letterHelper(digits, res, s+string(c), depth+1)
		}
		return *res
	}
	return letterHelper(digits, &[]string{}, "", 0)
}

//leetcode submit region end(Prohibit modification and deletion)
//给你一个字符串 S、一个字符串 T，请在字符串 S 里面找出：包含 T 所有字符的最小子串。
//
// 示例：
//
// 输入: S = "ADOBECODEBANC", T = "ABC"
//输出: "BANC"
//
// 说明：
//
//
// 如果 S 中不存这样的子串，则返回空字符串 ""。
// 如果 S 中存在这样的子串，我们保证它是唯一的答案。
//
// Related Topics 哈希表 双指针 字符串 Sliding Window

//leetcode submit region begin(Prohibit modification and deletion)
func minWindow(s string, t string) string {
	sLen, tLen := len(s), len(t)
	if sLen == 0 || tLen == 0 || sLen < tLen {
		return ""
	}
	var i, j, start, found int
	minLen, tCount, sCount := 0x7fffffff, [256]int{}, [256]int{}
	for _, c := range t {
		tCount[c]++
	}
	for j < sLen {
		if found < tLen {
			prev := int(s[j])
			if tCount[prev] > 0 {
				sCount[prev]++
				if sCount[prev] <= tCount[prev] {
					found++
				}
			}
			j++
		}
		for found == tLen {
			if j-i < minLen {
				start, minLen = i, j-i
			}
			next := int(s[i])
			if tCount[next] > 0 {
				sCount[next]--
				if sCount[next] < tCount[next] {
					found--
				}
			}
			i++
		}
	}
	if minLen == 0x7fffffff {
		return ""
	}
	return s[start : start+minLen]
}

//leetcode submit region end(Prohibit modification and deletion)

//给你一个整数数组 nums，请你将该数组升序排列。
//
//
//
//
//
//
// 示例 1：
//
// 输入：nums = [5,2,3,1]
//输出：[1,2,3,5]
//
//
// 示例 2：
//
// 输入：nums = [5,1,1,2,0,0]
//输出：[0,0,1,1,2,5]
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 50000
// -50000 <= nums[i] <= 50000
//
//

//leetcode submit region begin(Prohibit modification and deletion)
func sortArray(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	partition := func(nums []int, begin int, end int) int {
		pivot, counter := end, begin
		for i := begin; i < end; i++ {
			if nums[i] < nums[pivot] {
				nums[i], nums[counter] = nums[counter], nums[i]
				counter++
			}
		}
		nums[counter], nums[pivot] = nums[pivot], nums[counter]
		return counter
	}
	var quickSort func(nums []int, begin int, end int)
	quickSort = func(nums []int, begin int, end int) {
		if begin < end {
			pivot := partition(nums, begin, end)
			quickSort(nums, begin, pivot-1)
			quickSort(nums, pivot+1, end)
		}
	}
	quickSort(nums, 0, len(nums)-1)
	return nums
}

//leetcode submit region end(Prohibit modification and deletion)

//给定一个 N 叉树，返回其节点值的层序遍历。 (即从左到右，逐层遍历)。
//
// 例如，给定一个 3叉树 :
//
//
//
//
//
//
//
// 返回其层序遍历:
//
// [
//     [1],
//     [3,2,4],
//     [5,6]
//]
//
//
//
//
// 说明:
//
//
// 树的深度不会超过 1000。
// 树的节点总数不会超过 5000。
// Related Topics 树 广度优先搜索

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */
type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	queue := []*Node{root}
	for len(queue) > 0 {
		var tmp []int
		var next []*Node
		for _, q := range queue {
			tmp = append(tmp, q.Val)
			if q.Children != nil && len(q.Children) > 0 {
				next = append(next, q.Children...)
			}
		}
		res, queue = append(res, tmp), next
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
//给定一个 N 叉树，返回其节点值的后序遍历。
//
// 例如，给定一个 3叉树 :
//
//
//
//
//
//
//
// 返回其后序遍历: [5,6,3,2,4,1].
//
//
//
// 说明: 递归法很简单，你可以使用迭代法完成此题吗? Related Topics 树

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func postorder(root *Node) []int {
	if root == nil {
		return []int{}
	}
	var res []int
	var helper func(node *Node)
	helper = func(node *Node) {
		if node == nil {
			return
		}
		if node.Children != nil {
			for _, val := range node.Children {
				helper(val)
			}
		}
		res = append(res, node.Val)
	}
	helper(root)
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
//一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。
//
// 机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。
//
// 现在考虑网格中有障碍物。那么从左上角到右下角将会有多少条不同的路径？
//
//
//
// 网格中的障碍物和空位置分别用 1 和 0 来表示。
//
// 说明：m 和 n 的值均不超过 100。
//
// 示例 1:
//
// 输入:
//[
//  [0,0,0],
//  [0,1,0],
//  [0,0,0]
//]
//输出: 2
//解释:
//3x3 网格的正中间有一个障碍物。
//从左上角到右下角一共有 2 条不同的路径：
//1. 向右 -> 向右 -> 向下 -> 向下
//2. 向下 -> 向下 -> 向右 -> 向右
//
// Related Topics 数组 动态规划

//leetcode submit region begin(Prohibit modification and deletion)
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if obstacleGrid == nil || len(obstacleGrid) == 0 {
		return 0
	}
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	if (obstacleGrid[0][0] | obstacleGrid[m-1][n-1]) == 1 {
		return 0
	}
	steps := getXYArray(m, n)
	steps[0][0] = 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				steps[i][j] = 0
			} else if i == 0 {
				if j > 0 {
					steps[i][j] = steps[i][j-1]
				}
			} else if j == 0 {
				steps[i][j] = steps[i-1][j]
			} else {
				steps[i][j] = steps[i][j-1] + steps[i-1][j]
			}
		}
	}
	return steps[m-1][n-1]
}

func getXYArray(m int, n int) [][]int {
	var steps [][]int
	for i := 0; i < m; i++ {
		var temp []int
		for j := 0; j < n; j++ {
			temp = append(temp, 0)
		}
		steps = append(steps, temp)
	}
	return steps
}

//leetcode submit region end(Prohibit modification and deletion)
//你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上
//被小偷闯入，系统会自动报警。
//
// 给定一个代表每个房屋存放金额的非负整数数组，计算你在不触动警报装置的情况下，能够偷窃到的最高金额。
//
// 示例 1:
//
// 输入: [1,2,3,1]
//输出: 4
//解释: 偷窃 1 号房屋 (金额 = 1) ，然后偷窃 3 号房屋 (金额 = 3)。
//     偷窃到的最高金额 = 1 + 3 = 4 。
//
// 示例 2:
//
// 输入: [2,7,9,3,1]
//输出: 12
//解释: 偷窃 1 号房屋 (金额 = 2), 偷窃 3 号房屋 (金额 = 9)，接着偷窃 5 号房屋 (金额 = 1)。
//     偷窃到的最高金额 = 2 + 9 + 1 = 12 。
//
// Related Topics 动态规划

//leetcode submit region begin(Prohibit modification and deletion)
func rob(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	n, robStatus := len(nums), [][2]int{{0, nums[0]}}
	for i := 1; i < n; i++ {
		robStatus = append(robStatus, [2]int{0, 0})
		robStatus[i][0] = int(math.Max(float64(robStatus[i-1][0]), float64(robStatus[i-1][1])))
		robStatus[i][1] = nums[i] + robStatus[i-1][0]
	}
	return int(math.Max(float64(robStatus[n-1][0]), float64(robStatus[n-1][1])))
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	/*node := &Node{
		Val: 1,
		Children: []*Node{
			&Node{Val: 2, Children: []*Node{}},
			&Node{Val: 3, Children: []*Node{}},
		},
	}*/
	fmt.Println(rob([]int{2, 7, 9, 3, 1}))
}
