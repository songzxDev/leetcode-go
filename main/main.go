package main

import (
	"fmt"
	"sort"
)
func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	if nums != nil && len(nums) > 2 {
		sort.Ints(nums)
		n, first := len(nums), nums[0]
		for i := 0; first <= 0 && i < n-2; i++ {
			if i == 0 || nums[i] > nums[i-1] {
				j, k := i+1, n-1
				for j < k {
					numAdd := nums[i] + nums[j] + nums[k]
					if numAdd == 0 {
						res = append(res, []int{nums[i], nums[j], nums[k]})
						j++
						k--
						for j < k && nums[j] == nums[j-1] {
							j++
						}
						for j < k && nums[k] == nums[k + 1] {
							k--
						}
					} else if numAdd < 0 {
						j++
					} else {
						k--
					}
				}
			}
		}
	}
	return res
}
func main() {
	/*abc := week01.ladderLength("hit",
		"cog",
		[]string{"hot", "dot", "dog", "lot", "log", "cog"})
	fmt.Println(abc)*/
	/*fmt.Println([]rune("abc"))
	fmt.Println(string(97))
	fmt.Println(string([]rune{97, 98, 99}))

	a := []int{1, 0, 9, -892, 267}
	sort.Ints(a)
	fmt.Println(a)*/
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}
