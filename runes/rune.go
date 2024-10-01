package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const str string = "\x3d\x20\xe2\x8c\x98"
	fmt.Println(str)
	fmt.Printf("%+q\n", str)
	var character rune = 20000
	fmt.Printf("%c\n", character)
	b := []byte("Hello, 世界")

	for len(b) > 0 {
		r, size := utf8.DecodeRune(b)
		fmt.Printf("%c %v\n", r, size)
		b = b[size:]
	}
}
