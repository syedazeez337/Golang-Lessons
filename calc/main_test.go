package main

import (
	"math/rand"
	"testing"
)

// go test finds for the file *_test.go
/*
func TestAdd(t *testing.T) {
	result := Add(2, 3)
	expected := 5
	if result != expected {
		t.Errorf("Add(2,3) = %d; want %d", result, expected)
	}
}

func TestSub(t *testing.T) {
	result := Sub(5, 2)
	expected := 3
	if result != expected {
		t.Errorf("Sub(5,2) = %d; want %d", result, expected)
	}
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(1, 2)
	}
}
*/

func BenchmarkRandInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rand.Int()
	}
}
