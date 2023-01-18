package main

import (
	"fmt"
	"testing"
)

// func TestMySum(t *testing.T) {
// 	x := mySum(2, 3)
// 	if x != 5 {
// 		t.Error("Expected 5, got", x)
// 	}
// }

// go test ./...
// go -http=:8080 > examples
func ExampleMySum() {
	fmt.Println(MySum(2, 3))
	// Output:
	// 5
}

// go test
func TestMySum(t *testing.T) {
	type test struct {
		Data   []int
		Answer int
	}

	tests := []test{
		test{[]int{1, 2, 3}, 6},
		test{[]int{1, 2, -1}, 2},
		test{[]int{1, -2, -3}, -4},
		test{[]int{1, 0, -1}, 0},
		test{[]int{3, 3, 3}, 9},
	}

	for _, v := range tests {
		x := MySum(v.Data...)
		if x != v.Answer {
			t.Error("Expected ", v.Answer, "Got ", x)
		}
	}
}

// go test -bench .
func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MySum(i, i+1)
	}
}
