// https://leetcode.com/problems/group-anagrams/


func groupAnagrams(strs []string) [][]string {
	res := make([][]string, 0, len(strs))
	if len(strs) == 0 {
		return res
	}
	m := make(map[[26]uint8][]string)
	flag := int('a')
	for _, str := range strs {
		var arr [26]uint8
		for _, char := range str {
			arr[int(char)-flag]++
		}
		m[arr] = append(m[arr], str)
	}
	for _, v := range m {
		res = append(res, v)
	}
	return res
}

//
//func groupAnagrams(strs []string) [][]string {
//	var result [][]string
//	m := make(map[int][]string)
//	for i := range strs {
//		key := 1
//		for j := range strs[i] {
//			key += int(strs[i][j])
//		}
//		for j := range strs[i] {
//			key *= int(strs[i][j])
//		}
//		m[key] = append(m[key], strs[i])
//	}
//	for k := range m {
//		result = append(result, m[k])
//	}
//	return result
//}
