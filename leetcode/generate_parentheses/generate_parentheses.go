// https://leetcode.com/problems/generate-parentheses/
package main

func generateParenthesis(n int) []string {
	var rv []string
	rv = backtrack("", rv, n, 0, 0)
	return rv
}

func backtrack(rs string, rv []string, n, left, right int) []string {
	// https://en.wikipedia.org/wiki/Catalan_number
	if len(rs) == n*2 {
		rv = append(rv, rs)
		return rv
	}
	if left < n {
		rs += "("
		rv = backtrack(rs, rv, n, left+1, right)
		rs = rs[0 : len(rs)-1]
	}
	if right < left {
		rs += ")"
		rv = backtrack(rs, rv, n, left, right+1)
		rs = rs[0 : len(rs)-1]
	}
	return rv
}
