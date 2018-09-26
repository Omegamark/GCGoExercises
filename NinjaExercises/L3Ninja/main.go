package main

import "fmt"

func main() {
	// exercise5()
	// exercise6("antidisestablismentarianism")
	// exercise7()
	// exercise8()
	// exercise9()
	exercise10()
}

func exercise5() {
	for i := 0; i <= 100; i++ {
		fmt.Printf("When %v is divided by 4, the remainder (aka modulus) is %v\n", i, i%4)
	}
}

func exercise6(c string) {
	if len(c) > 10 {
		fmt.Println("That's a long word...")
	}
}

func exercise7() {
	x := "antidisestablismentarianism"
	if x == "" {
		fmt.Println("please enter a word")
	} else if len(x) > len(x)+1 {
		fmt.Println("there is no word in the english language that long")
	} else {
		fmt.Println("doody")
	}
}

func exercise8() {
	switch {
	case false:
		fmt.Println("should not print")
	case true:
		fmt.Println("should print")
	}
}

func exercise9() {
	favSport := "Capoeira"
	switch favSport {
	case "skiing":
		fmt.Println("well, I'm in Colorado")
	case "swimming":
		fmt.Println("Not for fun")
	case "surfing":
		fmt.Println("Time to move to Hawaii")
	case "Capoeira":
		fmt.Println("Wish I had time again")
	}
}

func exercise10() {
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)
}
