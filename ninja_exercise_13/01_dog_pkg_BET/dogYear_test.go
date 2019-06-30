package dog1

import (
	"fmt"
	"testing"
)

var humanYear = 5

func TestYears(t *testing.T) {
	x := Years(humanYear)
	if x != 35 {
		t.Error("Expected :", 35, "Got :", x)
	}
}

func ExampleYears() {
	fmt.Println(Years(5))
	//Output:
	// 35

}

func BenchmarkDogYear(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(humanYear)
	}
}
