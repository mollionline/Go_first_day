package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	banner("Go", 6)
	banner("G☺", 6)
	fmt.Println("g", isPolindrome("g"))
	fmt.Println("go", isPolindrome("go"))
	fmt.Println("gog", isPolindrome("gog"))
	fmt.Println("g☺g", isPolindrome("g☺g"))

	s := "G☺"
	fmt.Println("len", len(s))

	for i, r := range s {
		fmt.Println(i, r)
		if i == 0 {
			fmt.Printf("%c of type %T\n", r, r)
		}
	}

	x, y := 1, "1"
	fmt.Printf("x=%v, y=%v\n", x, y)
	fmt.Printf("x=%#v, y=%#v\n", x, y)

	// byte (uit8)
	// rune (int32)
	b := s[0]
	c := s[1]
	fmt.Printf("%c of type %T\n", b, b)
	fmt.Printf("%c of type %T\n", c, c)
}

func isPolindrome(s string) bool {
	// TODO: Your code goes here
	rs := []rune(s) // get slice of runes out of s
	for i := 0; i < len(rs)/2; i++ {
		if rs[i] != rs[len(rs)-1] {
			return false
		}
	}
	return true
}

func banner(text string, width int) {
	padding := (width - utf8.RuneCountInString(text)) / 2
	// padding := (width - len(text)) / 2

	for i := 0; i < padding; i++ {
		fmt.Print(" ")

	}
	fmt.Println(text)
	for i := 0; i < width; i++ {
		fmt.Print("-")
	}
	fmt.Println()
}
