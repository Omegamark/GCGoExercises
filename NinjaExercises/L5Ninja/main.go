package main

import "fmt"

//exported Person stuct
type Person struct {
	FirstName   string
	LastName    string
	FavIceCream []string
}

// Non exported vehicle struct
type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	turbo bool
}

func main() {
	// exercise1()
	// exercise2()
	// exercise3()
	exercise4()
}

func exercise1() {
	p1 := Person{
		FirstName: "Mark",
		LastName:  "Grant",
		FavIceCream: []string{
			"The Americone Dream",
			"Caramel",
			"Mint Chocolate",
		},
	}

	p2 := Person{
		FirstName: "Smoove",
		LastName:  "McGroove",
		FavIceCream: []string{
			"Nutty Buddy",
			"Vanilla",
			"Bubblegum",
		},
	}

	for _, v := range p1.FavIceCream {
		fmt.Println(v)
	}

	for _, v := range p2.FavIceCream {
		fmt.Printf("I %v %v Love to eat: %v\n", p2.FirstName, p2.LastName, v)
	}
}

func exercise2() {
	p1 := Person{
		FirstName: "Mark",
		LastName:  "Grant",
		FavIceCream: []string{
			"The Americone Dream",
			"Caramel",
			"Mint Chocolate",
		},
	}

	p2 := Person{
		"Smoove",
		"McGroove",
		[]string{
			"Nutty Buddy",
			"Vanilla",
			"Bubblegum",
		},
	}

	m := map[string]Person{
		p1.LastName: p1,
		p2.LastName: p2,
	}

	fmt.Println(m)

	for _, v := range m {

		fmt.Println("---------------------")
		fmt.Println(v.FirstName)
		fmt.Println(v.LastName)
		for i, val := range v.FavIceCream {
			fmt.Println(i, val)
		}
		fmt.Println("---------------------")
	}
}

func exercise3() {
	t1 := truck{
		vehicle{
			2,
			"blue",
		},
		true,
	}

	s1 := sedan{
		vehicle{
			2,
			"black",
		},
		true,
	}

	fmt.Println(t1, s1)
	fmt.Println(t1.doors)
	fmt.Println(s1.turbo)

}

func exercise4() {
	anonStruct := struct {
		name     string
		friends  map[string]int
		favGames []string
	}{
		"Mark",
		map[string]int{
			"Tyler": 7205550000,
			"Iggy":  4305551111,
		},
		[]string{"Shadow of the Colossus", "Ratchet & Clank", "Zelda", "Alundra"},
	}

	fmt.Println(anonStruct)
}
