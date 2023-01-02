package main

import "testing"

// Experiment to measure the difference in running time between our potentially inefficient versions and the one using `strings.Join()`.

func BenchmarkIsPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {

	}
}

func IsPalindrome() {

}
