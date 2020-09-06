// https://leetcode.com/problems/count-primes/
package main

import (
	"math"
)

//func countPrimes(n int) int {
//	if n < 3 {
//		return 0
//	}
//	primes := make([]int, n)
//	for i := 2; i < int(math.Sqrt(float64(n)))+1; i++ {
//		for j := i*i; j < n; {
//			primes[j] = 1
//			j += i
//		}
//	}
//	count := 0
//	for i := 2; i < n; i++ {
//		if primes[i] == 0 {
//			count++
//		}
//	}
//	return count
//}

// better
// https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes
func countPrimes(n int) int {
	if n < 3 {
		return 0
	}
	notPrimes := make([]bool, n)
	count := 0
	for i := 2; i < int(math.Sqrt(float64(n)))+1; i++ {
		if notPrimes[i] {
			continue
		}
		for j := 2; i*j < n; j++ {
			notPrimes[i*j] = true
		}
	}
	for i := 2; i < n; i++ {
		if !notPrimes[i] {
			count++
		}
	}
	return count
}
