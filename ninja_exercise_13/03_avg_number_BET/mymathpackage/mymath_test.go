package mymathpackage

import (
	"fmt"
	"testing"
)

func ExampleCenteredAvg() {
	fmt.Println(CenteredAvg([]int{5, 4, 3, 2, 1}))
	//Output:
	//3
}
func BencharkCenteredAvg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CenteredAvg([]int{5, 4, 3, 2, 1, 23323, 5000000})

	}
}

// func TestCenteredAvg(t *testing.T) {
// 	xi := []int{5, 4, 3, 2, 1}
// 	n := CenteredAvg(xi)

// 	if n != 2{
// 		t.Error("Expected:", 2, "Got:", n)
// 	}

func TestCoveredAvg(t *testing.T) {
	type test struct {
		data []int
		ans  float64
	}
	tests := []test{
		test{[]int{10, 20, 40, 60, 80}, 40},//60+40+20/3
		test{[]int{2, 4, 6, 8, 10, 12}, 7},
		test{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 5},
	}
	for _, v := range tests {
		f := CenteredAvg(v.data)
		if f != v.ans {
			t.Error("Got:", f, "Expected:", v.ans)
		}
	}
}
