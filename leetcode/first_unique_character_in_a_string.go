// https://leetcode.com/problems/first-unique-character-in-a-string/


func firstUniqChar(s string) int {
	if len(s) == 0 {return -1}
	letter := [26]int{}
	for _, v := range s {
		letter[v-97]++
	}
	for i, v := range s {
		if letter[v-97] == 1 {
			return i
		}
	}
	return -1
}
