//Find element exists in sorted two-dimensional array

package main

import (
	"fmt"
)

var arr = [4][4]int{
	{1, 2, 8, 9},
	{2, 4, 9, 12},
	{4, 7, 10, 13},
	{6, 8, 11, 15}}

func main() {
	fmt.Println(findNum(7, 0, 3, arr))
	fmt.Println(findNum(42, 0, 3, arr))
}

func findNum(target, n, m int, arr [4][4]int) bool {
	if n > 3 || m < 0 {
		return false
	}
	if target == arr[n][m] {
		return true
	} else if target < arr[n][m] {
		return findNum(target, n, m-1, arr)
	} else {
		return findNum(target, n+1, m, arr)
	}
}
