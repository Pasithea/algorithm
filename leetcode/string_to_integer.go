// https://leetcode.com/problems/string-to-integer-atoi/


func myAtoi(str string) int {
	if len(str) == 0 {return 0}
	var sign byte = ' '
	rev := 0
	begin := false
	for i := 0; i < len(str); i++ {
		isInt := isInteger(str[i])
		if begin && !isInt {break}
		if isInt {
			rev = rev * 10 + int(str[i] - '0')
			begin = true
			if rev > 1 << 31 {break}
			continue
		}
		if str[i] == '-' || str[i] == '+' {
			if sign != ' ' {break}
			sign = str[i]
			begin = true
		} else if str[i] == ' ' {
			continue
		} else {break}
	}

	if sign == '-' {rev *= -1}
	if rev > 1 << 31 - 1 {
		rev = 1 << 31 - 1
	} else if rev < -1 << 31 {
		rev = -1 << 31
	}
	return rev
}

func isInteger(b byte) bool {
	if b >= '0' && b <= '9' {
		return true
	}
	return false
}
