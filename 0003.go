package main

/*
给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
 */

func lengthOfLongestSubstring(s string) int {
	start, longest := 0, 0
	loc := map[rune]int{}
	for i, c := range s {
		if index, ok := loc[c]; ok && index >= start {
			longest = max(i-start, longest)
			start = index+1
		}
		loc[c] = i
	}
	return max(len(s)-start, longest)
}

//func max(i, j int) int {
//	if i >= j {
//		return i
//	}
//	return j
//}

// input:
// "abcabcbb"
// "bbbbb"
// "pwwkew"
// " "
// "abba"