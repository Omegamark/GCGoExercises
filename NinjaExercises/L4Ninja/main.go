package main

import "fmt"

func main() {
	// exercise1()
	// exercise2()
	// exercise3()
	// exercise4()
	// exercise5()
	// exercise6()
	// exercise7()
	// exercise8()
	// exercise9()
	exercise10()
}

func exercise1() {
	x := [5]int{1, 2, 3, 4, 5}
	for _, v := range x {
		fmt.Printf("This item in the array is of TYPE %T\n", v)
	}
}

func exercise2() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, v := range x {
		fmt.Printf("This item in the array is of TYPE %T\n", v)
	}
	fmt.Printf("%T\n", x)
}

func exercise3() {
	intArray := [20]int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 44, 45, 46, 47, 48, 43, 44, 45, 46, 47}
	var a = intArray[0:5]
	var b = intArray[5:10]
	var c = intArray[10:15]
	var d = intArray[15:20]

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}

func exercise4() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	y := append(x, 52, 53, 54, 55)
	for i := 0; i < len(y); i++ {
		fmt.Println(y[i])
	}
	fmt.Printf("%v : this slice is of type %T\n", y, y)
}

func exercise5() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	var y []int
	y = append(x[:3], x[6:]...)
	fmt.Println("I'm y:", y)
	fmt.Println("I'm x:", x)
}

func exercise6() {
	states := make([]string, 50, 50)
	states = []string{` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`}
	// Length of slice
	fmt.Println(len(states))
	// Array capacity
	fmt.Println(cap(states))
	// Print values with index WITHOUT using range
	for i := 0; i < len(states); i++ {
		fmt.Printf("Index: %d\t State: %v\n", i, states[i])
	}
}

func exercise7() {
	jb := []string{"James", "Bond", "Shaken, not stirred"}
	mp := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	sa := [][]string{jb, mp}
	fmt.Println(sa)
	for i, sas := range sa {
		fmt.Println("Record:", i)
		for j, isa := range sas {
			fmt.Printf("\t Index: %d \t Value: %v\n", j, isa)
		}
	}
}

func exercise8() {
	// These 2 maps mean exactly the same thing.

	m := make(map[string][]string)
	m["bond_james"] = []string{"Shaken, not stirred", "Martinis", "Women"}
	m["moneypenny_miss"] = []string{"James Bond", "Literature", "Computer Science"}
	m["no_dr"] = []string{"Being evil", "Ice cream", "Sunsets"}

	// m := map[string][]string{
	// 	`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
	// 	`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
	// 	`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	// }

	// Print the values along with their index position in the slice.
	for i, val := range m {
		fmt.Printf("These are %v's favorite things: %v\n", i, val)
	}
}

func exercise9() {
	m := make(map[string][]string)
	m["bond_james"] = []string{"Shaken, not stirred", "Martinis", "Women"}
	m["moneypenny_miss"] = []string{"James Bond", "Literature", "Computer Science"}
	m["no_dr"] = []string{"Being evil", "Ice cream", "Sunsets"}
	m["Dr. Nefarious"] = []string{"Galactic Conquest", "Screaming", "Kittens"}

	// Print the values along with their index position in the slice.
	for i, val := range m {
		fmt.Printf("These are %v's favorite things: %v\n", i, val)
	}
}

func exercise10() {
	m := make(map[string][]string)
	m["bond_james"] = []string{"Shaken, not stirred", "Martinis", "Women"}
	m["moneypenny_miss"] = []string{"James Bond", "Literature", "Computer Science"}
	m["no_dr"] = []string{"Being evil", "Ice cream", "Sunsets"}
	m["Dr. Nefarious"] = []string{"Galactic Conquest", "Screaming", "Kittens"}

	// Print the values along with their index position in the slice.
	for i, val := range m {
		fmt.Printf("These are %v's favorite things: %v\n", i, val)
	}
	delete(m, "no_dr")
	for i, val := range m {
		fmt.Printf("These are %v's favorite things: %v\n", i, val)
	}
}
