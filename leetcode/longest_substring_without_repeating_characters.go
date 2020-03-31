// https://leetcode.com/problems/longest-substring-without-repeating-characters/


func lengthOfLongestSubstring(s string) int {
	var (
		arr    [128]int
		maxLen int
		curLen int
	)
	for i, j := 0, 0; j < len(s); j++ {
		if arr[s[j]] > i {
			i = arr[s[j]]
		}
		curLen = j - i + 1
		if curLen > maxLen {
			maxLen = curLen
		}
		arr[s[j]] = j + 1
	}
	return maxLen
}

//func lengthOfLongestSubstring(s string) int {
//	var arr [128]int
//	b := []byte(s)
//	var maxLen, n int
//	for i, j := 0, 0; i < len(s) && j < len(s); j++ {
//		if arr[b[j]] > i {
//			i = arr[b[j]]
//		}
//		n = j - i + 1
//		if n > maxLen {
//			maxLen = n
//		}
//		arr[b[j]] = j + 1
//	}
//	return maxLen
//}
