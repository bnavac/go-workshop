package main

import (
	"errors"
	"fmt"
	//"unsafe"
)

func main() {
	fmt.Printf("Add: %d", add(2, 3))
	val, err := div(5, 0)
	if err != nil {
		fmt.Printf("%d, %s\n", val, err.Error())
	}
	var p *int = &val
	*p = 21
	fmt.Println(*p)
	var arr = make([]int, 12)
	//*(arr+3) = 2
	arr[3] = 2
	//up := unsafe.Pointer(&arr[0])
	//fmt.Println(*(*int)(unsafe.Pointer(uintptr(up) + uintptr(2))))
	fmt.Printf("Function as a paremter: %d", compute(add))
	var i func() int = func() int { return 5 }
	fmt.Printf("Function literal (anonymous function) %d", i())
	fmt.Println("closures")
	int1 := intSeq()
	int2 := intSeq()
	int1()
	fmt.Println(int1(), int2())

}
func add(x int, y int) int {
	return x + y
}
func div(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("Division by 0")
	}
	return x / y, nil
}
func compute(fn func(int, int) int) int {
	return fn(3, 4)
}
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
