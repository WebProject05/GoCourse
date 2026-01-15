package main

import "testing"

func Add(a int, b int) int {
	return a + b
}

// UNIT TEST (COMMENTED OUT)
//
// func TestAdd(t *testing.T) {
// 	result := Add(2, 3)
// 	expected := 5
// 	if result != expected {
// 		t.Errorf("Add(2, 3) = %d; want %d", result, expected)
// 	}
// }

// BENCHMARK

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(2, 3)
	}
}
