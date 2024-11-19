package util

import "testing"

// Unit test for IsPrime using table-driven tests
func TestIsPrime(t *testing.T) {
	tests := []struct {
		input    int
		expected bool
	}{
		{2, true},
		{4, false},
		{17, true},
		{18, false},
		{1, false},
	}

	for _, test := range tests {
		result := IsPrime(test.input)
		if result != test.expected {
			t.Errorf("IsPrime(%d) = %v; expected %v", test.input, result, test.expected)
		}
	}
}

// Unit test for ReverseString
func TestReverseString(t *testing.T) {
	result := ReverseString("hello")
	expected := "olleh"
	if result != expected {
		t.Errorf("ReverseString('hello') = %s; expected %s", result, expected)
	}
}

// Unit test for SumOfSlice
func TestSumOfSlice(t *testing.T) {
	result := SumOfSlice([]int{1, 2, 3, 4})
	expected := 10
	if result != expected {
		t.Errorf("SumOfSlice([1, 2, 3, 4]) = %d; expected %d", result, expected)
	}
}

// Unit test for Factorial
func TestFactorial(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{5, 120},
		{0, 1},
		{-1, -1},
		{1, 1},
	}

	for _, test := range tests {
		result := Factorial(test.input)
		if result != test.expected {
			t.Errorf("Factorial(%d) = %d; expected %d", test.input, result, test.expected)
		}
	}
}

// Benchmark for ReverseString
func BenchmarkReverseString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReverseString("hello world")
	}
}

// Benchmark for IsPrime
func BenchmarkIsPrime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPrime(101)
	}
}

// Benchmark for SumOfSlice
func BenchmarkSumOfSlice(b *testing.B) {
	slice := make([]int, 1000)
	for i := range slice {
		slice[i] = i
	}
	for i := 0; i < b.N; i++ {
		SumOfSlice(slice)
	}
}
