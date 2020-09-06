// https://leetcode.com/problems/fizz-buzz/
package main

import (
	"strconv"
)

func fizzBuzz(n int) []string {
	var rv []string
	for i := 1; i <= n; i++ {
		s := ""
		if i%3 == 0 {
			s += "Fizz"
			if i%5 == 0 {
				s += "Buzz"
			}
		} else if i%5 == 0 {
			s += "Buzz"
		} else {
			s = strconv.Itoa(i)
		}
		rv = append(rv, s)
	}
	return rv
}
