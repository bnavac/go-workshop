package main

import (
	"fmt"
	"math/cmplx"
)

func main() {
	var num int //which will be either int32 or int64 depending on pc.
	var posNum uint16
	var str string
	var False bool
	var character byte
	num = 3000
	fmt.Printf("Int: %T %v %c\n", num, num, num)
	//posNum = -15
	posNum = 15
	fmt.Printf("Uint: %T %v\n", posNum, posNum)
	str = "hello"
	fmt.Printf("String: %T %v\n", str, str)
	False = true
	fmt.Printf("Bool: %t\n", False)
	character = 97
	fmt.Printf("byte: %T %v\n", character, character)
	const double = 12.5
	//double = 2.2
	fmt.Printf("float: %T %v\n", double, double)
	complex := cmplx.Sqrt(-5 + 12i)
	fmt.Printf("complex: %T %v\n", complex, complex)
	//empty()
	casting()
}

var i int
var f float64
var b bool
var s string

func empty() {
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
func casting() {
	var posNum int = 15
	var double float64 = 12.5
	//var res = posNum + double
	var res = double + float64(posNum)
	fmt.Printf("explicit casting: %T, %v\n", res, res)
	var character byte = 'a'
	var character2 uint8 = 30
	fmt.Printf("implicit casting: %T, %v\n", character+character2, character+character2)
	//complexAddition := cmplx.Sqrt(-5+12i) + complex128(posNum)
	//complexAddition := cmplx.Sqrt(-5+12i) + complex(float64(posNum), 0)
	//fmt.Printf("complex + real: %T %v\n", complexAddition, complexAddition)
}
