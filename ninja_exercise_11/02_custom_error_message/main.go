package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	p1 := person{
		First:   "Mayuresh",
		Last:    "Khalegaonkar",
		Sayings: []string{"good", "better", "best"},
	}

	bs, err := toJson(p1)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bs))
}

func toJson(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)
	if err != nil {
		return []byte{}, fmt.Errorf("Json was having some error %v", err)

	}
	return bs, nil
}
