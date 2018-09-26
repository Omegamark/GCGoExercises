package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {
	// Must be no breaks in the string literal for JSON.
	s := `[{"First":"Jamesz","Last":"Bond","Age":32},{"First":"Miss","Last":"MoneyPenny","Age":27}]`

	//  *** If data has a carraige return in it as below, when attempting to unmarshal, the string literal contains a '\n' for the carraige return and as a result, will not json.Unmarshal.
	// s := `[{"First":"Jamesz","Last":"Bond","Age":32},{"First":"Miss","Last
	//     ":"MoneyPenny","Age":27}]`

	bs := []byte(s)
	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs)

	// Same meaning, both initialize people of type []perosn
	// people := []person{}
	var people []person

	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(people)

	for i, v := range people {
		fmt.Println("\nPerson Number:", i)
		fmt.Println(v.First, v.Last, v.Age)
	}
}
