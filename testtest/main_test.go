package main

import (
	"testing"
	"fmt"
)

type squareTest struct {
	in, out int
}

var squareTests = []squareTest{
	squareTest{1, 1},
	squareTest{2, 4},
	squareTest{5, 25},
	squareTest{-2, 4},
}

// テスト.
func TestSqure( t *testing.T) {
	for _, st := range squareTests {
		v := Square(st.in)
		if v != st.out {
			t.Errorf("Square(%d) = %d, want %d.", st.in, v, st.out)
		}
	}
}

// ベンチマーク.
func BenchmarkSquare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Square(10)
	}
}

// Example系.
func ExampleSquare() {
	v := Square(11)
	fmt.Println(v)
	// Output: 121
}











