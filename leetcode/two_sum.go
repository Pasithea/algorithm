// https://leetcode.com/problems/two-sum/


//func twoSum(nums []int, target int) []int {
//	m := make(map[int]int, len(nums))
//	for i, v := range nums {
//		m[v] = i
//	}
//	for v, i := range m {
//		other := target - v
//		if p, ok := m[other]; ok && other != v {
//			return []int{i, p}
//		}
//		if other == v {
//			for other_i := range nums {
//				if nums[other_i] == other && other_i != m[v] {
//					return []int{other_i, m[v]}
//				}
//			}
//		}
//	}
//	return []int{}
//}


func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		n, ok := m[v]
		m[target-v] = i
		if ok {
			return []int{i, n}
		}
	}
	return []int{}
}
