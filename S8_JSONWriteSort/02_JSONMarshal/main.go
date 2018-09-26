package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Capitalizing the struct does not export the fields if the field names are lowercase. Field names must be upper case in order to be exported.
type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := person{
		"Jamesz",
		"Bond",
		32,
	}
	p2 := person{
		"Miss",
		"MoneyPenny",
		27,
	}
	people := []person{p1, p2}

	fmt.Println(people)

	bs, err := json.Marshal(people)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bs))

}
