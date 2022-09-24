package main

/*
给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。

你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。

你可以按任意顺序返回答案。
 */

func twoSum(nums []int, target int) []int {
	if len(nums) < 2 {
		return []int{}
	}

	tmp := map[int]int{}
	tmp[nums[0]] = 0
	for i := 1; i < len(nums); i++ {
		if index, ok := tmp[target-nums[i]]; ok {
			return []int{index, i}
		}
		tmp[nums[i]] = i
	}
	return []int{}
}

// input:
// []int{2,7,11,15}, 9
// []int{3,2,4}, 6
// []int{3,3}, 6