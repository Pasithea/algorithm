// https://leetcode.com/problems/reverse-string/


func reverseString(s []byte)  {
	length := len(s)
	if length <= 1 {return}
	p1, p2 := 0, length-1
	for ; p1 < p2 ; {
		s[p1], s[p2] = s[p2], s[p1]
		p1 += 1
		p2 -= 1
	}
}
