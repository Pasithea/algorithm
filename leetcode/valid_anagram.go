// https://leetcode.com/problems/valid-anagram/


func isAnagram(s string, t string) bool {
	if len(s) != len(t) {return false}
	m := make([]int, 26)
	for i := 0; i < len(s); i++ {
		m[s[i]-97]++
	}
	for i := 0; i < len(s); i++ {
		m[t[i]-97]--
		if m[t[i]-97] < 0 {return false}
	}
	return true
}
