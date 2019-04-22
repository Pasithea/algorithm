// https://leetcode.com/problems/number-of-1-bits/


//func hammingWeight(num uint32) int {
//	w := 0
//	var m uint32 = 1
//	for i := 0; i < 32; i++ {
//		if num & m != 0 {
//			w++
//		}
//		m <<= 1
//	}
//	return w
//}

// better
func hammingWeight(num uint32) int {
	w := 0
	for num > 0 {
		w++
		num &= num - 1
	}
	return w
}
