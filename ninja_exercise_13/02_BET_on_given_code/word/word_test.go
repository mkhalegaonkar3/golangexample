package word

import (
	"fmt"
	"testing"

	"github.com/mkhalegaonkar3/golangexample/ninja_exercise_13/02_BET_on_given_code/quote"
)

func TestCount(t *testing.T) {

	n := Count("one two three")
	if n != 3 {
		t.Error("Expected :", 3, "Got :", n)
	}

}
func TestUseCount(t *testing.T) {
	m := UseCount("one two three three two")
	for k, v := range m {
		switch k {
		case "one":
			if v != 1 {
				t.Error("Expected", 1, "Got", v)
			}
		case "two":

			if v != 2 {
				t.Error("Expected", 2, "Got", v)
			}
		case "three":

			if v != 2 {
				t.Error("Expected", 2, "Got", v)
			}
		}
	}
}
func ExampleCount() {

	fmt.Println(Count("Hi , i am mayuresh"))
	//Output:
	//5
}
func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(quote.May)
	}
}
func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount(quote.May)
	}
}
