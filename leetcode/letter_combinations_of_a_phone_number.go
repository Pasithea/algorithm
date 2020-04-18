// https://leetcode.com/problems/letter-combinations-of-a-phone-number/


var (
	m = map[string][]string{
		"2": {"a", "b", "c"},
		"3": {"d", "e", "f"},
		"4": {"g", "h", "i"},
		"5": {"j", "k", "l"},
		"6": {"m", "n", "o"},
		"7": {"p", "q", "r", "s"},
		"8": {"t", "u", "v"},
		"9": {"w", "x", "y", "z"},
	}
)

func letterCombinations(digits string) []string {
	var rs []string
	if len(digits) == 0 {
		return rs
	}
	rs = backtrack("", digits, rs)
	return rs
}

func backtrack(combine, digits string, rs []string) []string {
	if len(digits) == 0 {
		rs = append(rs, combine)
	} else {
		for _, v := range m[string(digits[0])] {
			rs = backtrack(combine + v, digits[1:], rs)
		}
	}
	return rs
}
