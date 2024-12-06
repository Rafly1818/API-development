package main

import "testing"

func TestMain(t *testing.T) {
	expected := "Hello, Go Modules!"
	result := "Hello, Go Modules!"

	if result != expected {
		t.Errorf("Hasil tidak sesuai: got %v, want %v", result, expected)
	}
}
