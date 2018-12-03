package word

import (
	"fmt"
	"testing"

	"../quote"
)

func ExampleCount() {
	s := "Hello World"
	fmt.Println(Count(s))
	// Output:
	// 2
}

func TestCount(t *testing.T) {
	n := Count("one two three")
	if n != 3 {
		t.Error("got", n, "want", 3)
	}
}

func TestUseCount(t *testing.T) {
	m := UseCount("Hi")
	fmt.Println("I should be a map:", m)
	if m == nil {
		t.Error("got", m, "want", map[string]int{"Hi": 1})
	}
	m = UseCount("one two three")
	for k, v := range m {
		switch k {
		case "one":
			if v != 1 {
				t.Error("got", v, "want", 1)
			}
		case "two":
			if v != 1 {
				t.Error("got", v, "want", 1)
			}
		case "three":
			if v != 1 {
				t.Error("got", v, "want", 1)
			}
		}
	}
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(quote.SunAlso)
	}
}

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount(quote.SunAlso)
	}
}
