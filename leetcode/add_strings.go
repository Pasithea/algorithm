// https://leetcode.com/problems/add-strings/


func addStrings(num1 string, num2 string) string {
	carry := 0
	rs := ""
	for len(num1) != 0 || len(num2) != 0 || carry != 0 {
		var (
			t1 int
			t2 int
			t  int
		)
		if len(num1) != 0 {
			t1, _ = strconv.Atoi(string(num1[len(num1)-1]))
			num1 = num1[:len(num1)-1]
		}
		if len(num2) != 0 {
			t2, _ = strconv.Atoi(string(num2[len(num2)-1]))
			num2 = num2[:len(num2)-1]
		}
		t = t1 + t2 + carry
		carry = t / 10
		if carry != 0 {
			t %= 10
		}
		ts := strconv.Itoa(t)
		rs = ts + rs
	}
	return rs
}
