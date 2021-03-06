package main

import (
	"fmt"
	"math"
	"sort"
)

//给定一个排序数组，你需要在 原地 删除重复出现的元素，使得每个元素只出现一次，返回移除后数组的新长度。
//
// 不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。
//
//
//
// 示例 1:
//
// 给定数组 nums = [1,1,2],
//
//函数应该返回新的长度 2, 并且原数组 nums 的前两个元素被修改为 1, 2。
//
//你不需要考虑数组中超出新长度后面的元素。
//
// 示例 2:
//
// 给定 nums = [0,0,1,1,1,2,2,3,3,4],
//
//函数应该返回新的长度 5, 并且原数组 nums 的前五个元素被修改为 0, 1, 2, 3, 4。
//
//你不需要考虑数组中超出新长度后面的元素。
//
//
//
//
// 说明:
//
// 为什么返回数值是整数，但输出的答案是数组呢?
//
// 请注意，输入数组是以「引用」方式传递的，这意味着在函数里修改输入数组对于调用者是可见的。
//
// 你可以想象内部操作如下:
//
// // nums 是以“引用”方式传递的。也就是说，不对实参做任何拷贝
//int len = removeDuplicates(nums);
//
//// 在函数里修改输入数组对于调用者是可见的。
//// 根据你的函数返回的长度, 它会打印出数组中该长度范围内的所有元素。
//for (int i = 0; i < len; i++) {
//    print(nums[i]);
//}
//
// Related Topics 数组 双指针

//leetcode submit region begin(Prohibit modification and deletion)
func removeDuplicates(nums []int) int {
	r := 0
	for i := 0; i < len(nums); i++ {
		if i == 0 || nums[i-1] < nums[i] {
			nums[r] = nums[i]
			r++
		}
	}
	return r
}

//leetcode submit region end(Prohibit modification and deletion)
//将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
//
//
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
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var res []int
	var treeHelper func(node *TreeNode)
	treeHelper = func(node *TreeNode) {
		if node == nil {
			return
		}
		treeHelper(node.Left)
		res = append(res, node.Val)
		treeHelper(node.Right)
	}
	treeHelper(root)
	return res
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
	var res []int
	var treeHelper func(node *TreeNode)
	treeHelper = func(node *TreeNode) {
		if node == nil {
			return
		}
		res = append(res, node.Val)
		treeHelper(node.Left)
		treeHelper(node.Right)
	}
	treeHelper(root)
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
//根据一棵树的前序遍历与中序遍历构造二叉树。
//
// 注意:
//你可以假设树中没有重复的元素。
//
// 例如，给出
//
// 前序遍历 preorder = [3,9,20,15,7]
//中序遍历 inorder = [9,3,15,20,7]
//
// 返回如下的二叉树：
//
//     3
//   / \
//  9  20
//    /  \
//   15   7
// Related Topics 树 深度优先搜索 数组

//leetcode submit region begin(Prohibit modification and deletion)

func buildTree(preorder []int, inorder []int) *TreeNode {
	inorderMap := make(map[int]int, len(inorder))
	for i, n := range inorder {
		inorderMap[n] = i
	}
	var buildHelper func(preLeft, preRight, inLeft, inRight int) *TreeNode
	buildHelper = func(preLeft, preRight, inLeft, inRight int) *TreeNode {
		if preLeft > preRight || inLeft > inRight {
			return nil
		}
		inRootIdx := inorderMap[preorder[preLeft]]
		leftSubtreeLen := inRootIdx - 1 - inLeft + 1
		return &TreeNode{
			Val:   preorder[preLeft],
			Left:  buildHelper(preLeft+1, preLeft+leftSubtreeLen, inLeft, inRootIdx-1),
			Right: buildHelper(preLeft+leftSubtreeLen+1, preRight, inRootIdx+1, inRight),
		}
	}
	return buildHelper(0, len(preorder)-1, 0, len(inorder)-1)
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
	primes := [26]int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101}
	getNumKey, groupMap := func(stt string) int {
		numKey := 1
		for _, c := range stt {
			numKey *= primes[c-'a']
		}
		return numKey
	}, make(map[int][]string)
	for _, stt := range strs {
		numKey := getNumKey(stt)
		groupMap[numKey] = append(groupMap[numKey], stt)
	}
	for k := range groupMap {
		res = append(res, groupMap[k])
	}
	return res
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
	sLen, tLen, i, j, start, minLen, found := len(s), len(t), 0, 0, 0, math.MaxInt32, 0
	if sLen < tLen {
		return ""
	}
	sMap, tMap := [256]int{}, [256]int{}
	for _, c := range t {
		tMap[c]++
	}
	for j < sLen {
		if found < tLen {
			before := s[j]
			if tMap[before] > 0 {
				sMap[before]++
				if sMap[before] <= tMap[before] {
					found++
				}
			}
			j++
		}
		for found == tLen {
			after := s[i]
			if j-i < minLen {
				start, minLen = i, j-i
			}
			if tMap[after] > 0 {
				sMap[after]--
				if sMap[after] < tMap[after] {
					found--
				}
			}
			i++
		}
	}
	if minLen == math.MaxInt32 {
		return ""
	}
	return s[start : start+minLen]
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
	isRun, n := nums[0] <= 0, len(nums)
	for i := 0; isRun && i < n-2; i++ {
		if i == 0 || nums[i-1] < nums[i] {
			j, k := i+1, n-1
			for j < k {
				add := nums[i] + nums[j] + nums[k]
				if add == 0 {
					res = append(res, []int{nums[i], nums[j], nums[k]})
					j, k = j+1, k-1
					for j < k && nums[j-1] == nums[j] {
						j++
					}
					for j < k && nums[k+1] == nums[k] {
						k--
					}
				} else if add < 0 {
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
//给定一个二叉树，判断其是否是一个有效的二叉搜索树。
//
// 假设一个二叉搜索树具有如下特征：
//
//
// 节点的左子树只包含小于当前节点的数。
// 节点的右子树只包含大于当前节点的数。
// 所有左子树和右子树自身必须也是二叉搜索树。
//
//
// 示例 1:
//
// 输入:
//    2
//   / \
//  1   3
//输出: true
//
//
// 示例 2:
//
// 输入:
//    5
//   / \
//  1   4
//     / \
//    3   6
//输出: false
//解释: 输入为: [5,1,4,null,null,3,6]。
//     根节点的值为 5 ，但是其右子节点值为 4 。
//
// Related Topics 树 深度优先搜索

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	var validHelper func(node *TreeNode, leftFar, rightFar int) bool
	validHelper = func(node *TreeNode, leftFar, rightFar int) bool {
		if node == nil {
			return true
		}
		if leftFar >= node.Val || node.Val >= rightFar {
			return false
		}
		return validHelper(node.Left, leftFar, node.Val) && validHelper(node.Right, node.Val, rightFar)
	}
	return validHelper(root, math.MinInt64, math.MaxInt64)
}

//leetcode submit region end(Prohibit modification and deletion)
//翻转一棵二叉树。
//
// 示例：
//
// 输入：
//
//      4
//   /   \
//  2     7
// / \   / \
//1   3 6   9
//
// 输出：
//
//      4
//   /   \
//  7     2
// / \   / \
//9   6 3   1
//
// 备注:
//这个问题是受到 Max Howell 的 原问题 启发的 ：
//
// 谷歌：我们90％的工程师使用您编写的软件(Homebrew)，但是您却无法在面试时在白板上写出翻转二叉树这道题，这太糟糕了。
// Related Topics 树

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left, right := root.Left, root.Right
	root.Left, root.Right = invertTree(right), invertTree(left)
	return root
}

//leetcode submit region end(Prohibit modification and deletion)
//给定一个二叉树，找出其最大深度。
//
// 二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
//
// 说明: 叶子节点是指没有子节点的节点。
//
// 示例：
//给定二叉树 [3,9,20,null,null,15,7]，
//
//     3
//   / \
//  9  20
//    /  \
//   15   7
//
// 返回它的最大深度 3 。
// Related Topics 树 深度优先搜索

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + int(math.Max(float64(maxDepth(root.Left)), float64(maxDepth(root.Right))))
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

type SortArrayUtil struct{}

func (sort *SortArrayUtil) partition(array []int, begin, end int) int {
	pivot, counter := end, begin
	for i := begin; i < end; i++ {
		if array[i] < array[pivot] {
			array[i], array[counter] = array[counter], array[i]
			counter++
		}
	}
	array[pivot], array[counter] = array[counter], array[pivot]
	return counter
}
func (sort *SortArrayUtil) QuickSort(array []int, begin, end int) {
	if begin >= end {
		return
	}
	pivot := sort.partition(array, begin, end)
	sort.QuickSort(array, begin, pivot-1)
	sort.QuickSort(array, pivot+1, end)
}

//leetcode submit region begin(Prohibit modification and deletion)
func sortArray(nums []int) []int {
	if nums != nil && len(nums) > 1 {
		var sortUtil SortArrayUtil
		sortUtil.QuickSort(nums, 0, len(nums)-1)
	}
	return nums
}

//leetcode submit region end(Prohibit modification and deletion)
//给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
//
// 岛屿总是被水包围，并且每座岛屿只能由水平方向或竖直方向上相邻的陆地连接形成。
//
// 此外，你可以假设该网格的四条边均被水包围。
//
//
//
// 示例 1:
//
// 输入:
//[
//['1','1','1','1','0'],
//['1','1','0','1','0'],
//['1','1','0','0','0'],
//['0','0','0','0','0']
//]
//输出: 1
//
//
// 示例 2:
//
// 输入:
//[
//['1','1','0','0','0'],
//['1','1','0','0','0'],
//['0','0','1','0','0'],
//['0','0','0','1','1']
//]
//输出: 3
//解释: 每座岛屿只能由水平和/或竖直方向上相邻的陆地连接而成。
//
// Related Topics 深度优先搜索 广度优先搜索 并查集

//leetcode submit region begin(Prohibit modification and deletion)
func numIslands(grid [][]byte) int {
	lands := 0
	var sink func(i, j int) int
	sink = func(i, j int) int {
		if i < 0 || j < 0 || i == len(grid) || j == len(grid[i]) || grid[i][j] == '0' {
			return 0
		}
		grid[i][j] = '0'
		sink(i-1, j)
		sink(i+1, j)
		sink(i, j-1)
		sink(i, j+1)
		return 1
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				lands += sink(i, j)
			}
		}
	}
	return lands
}

//leetcode submit region end(Prohibit modification and deletion)
//n 皇后问题研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
//
//
//
// 上图为 8 皇后问题的一种解法。
//
// 给定一个整数 n，返回所有不同的 n 皇后问题的解决方案。
//
// 每一种解法包含一个明确的 n 皇后问题的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。
//
// 示例:
//
// 输入: 4
//输出: [
// [".Q..",  // 解法 1
//  "...Q",
//  "Q...",
//  "..Q."],
//
// ["..Q.",  // 解法 2
//  "Q...",
//  "...Q",
//  ".Q.."]
//]
//解释: 4 皇后问题存在两个不同的解法。
//
//
//
//
// 提示：
//
//
// 皇后，是国际象棋中的棋子，意味着国王的妻子。皇后只做一件事，那就是“吃子”。当她遇见可以吃的棋子时，就迅速冲上去吃掉棋子。当然，她横、竖、斜都可走一到七步
//，可进可退。（引用自 百度百科 - 皇后 ）
//
// Related Topics 回溯算法

//leetcode submit region begin(Prohibit modification and deletion)
func solveNQueens(n int) [][]string {
	var res [][]string
	if n > 0 {
		var rows []int
		lie, pie, na := make(map[int]bool, n), make(map[int]bool, n), make(map[int]bool, n)
		toBoards := func(rows []int, n int) []string {
			var boards []string
			for _, r := range rows {
				var target []rune
				for i := 0; i < n; i++ {
					if r == i {
						target = append(target, 'Q')
					} else {
						target = append(target, '.')
					}
				}
				boards = append(boards, string(target))
			}
			return boards
		}
		containsKey := func(myMap map[int]bool, l int) bool {
			_, ok := myMap[l]
			return ok
		}
		var bfsHelper func(row, n int)
		bfsHelper = func(row, n int) {
			if row >= n {
				res = append(res, toBoards(rows, n))
				return
			}
			for l := 0; l < n; l++ {
				if containsKey(lie, l) || containsKey(pie, row+l) || containsKey(na, row-l) {
					continue
				}
				rows, lie[l], pie[row+l], na[row-l] = append(rows, l), true, true, true
				bfsHelper(row+1, n)
				rows = rows[:len(rows)-1]
				delete(lie, l)
				delete(pie, row+l)
				delete(na, row-l)
			}
		}
		bfsHelper(0, n)
	}
	return res
}
//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(solveNQueens(4))
}
