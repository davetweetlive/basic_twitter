package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := person{
		First: "DAve",
		Last:  "Augustus",
		Age:   27,
	}

	p2 := person{
		First: "Irak",
		Last:  "Rigia",
		Age:   27,
	}

	people := []person{p1, p2}
	fmt.Println(people)

	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println("Couldn't marshal JSON")
	}
	fmt.Println(string(bs))
}
