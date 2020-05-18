package week01

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
	if wordList == nil || len(wordList) == 0 {
		return 0
	}
	step := 1
	wordRune := []rune("abcdefghijklmnopqrstuvwxyz")
	beginSet, endSet, wordSet := getWordMap([]string{beginWord}), getWordMap([]string{endWord}), getWordMap(wordList)
	if _, ok := wordSet[endWord]; ok {
		for len(beginSet) > 0 {
			step++
			nextSet := make(map[string]bool)
			if len(beginSet) > len(endSet) {
				beginSet, endSet = endSet, beginSet
			}
			for word := range beginSet {
				for i := 0; i < len(word); i++ {
					for _, c := range wordRune {
						if c != int32(word[i]) {
							target := word[:i] + string(c) + word[i+1:]
							if _, ok := endSet[target]; ok {
								return step
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
	mySet := make(map[string]bool)
	for _, word := range wordList {
		mySet[word] = true
	}
	return mySet
}

//leetcode submit region end(Prohibit modification and deletion)
