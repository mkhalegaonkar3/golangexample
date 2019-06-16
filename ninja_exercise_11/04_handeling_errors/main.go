package main

import (
	"github.com/GoesToEleven/go-programming/code_samples/007-documentation/01/mymath"
	"fmt"
	"log"
)

type sqrtError struct {
	lat  string
	long string
	err  error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("math error: %v %v %v", se.lat, se.long, se.err)
}

func main() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}
mymath.Sum(5,3,4,5)

func sqrt(f float64) (float64, error) {
	if f < 0 {
		//e := errors.New("negative value error")
		e := fmt.Errorf("negative value error")
		return 0, sqrtError{"50.2289 N", "99.4656 W", e}
	}
	return 42, nil
}
