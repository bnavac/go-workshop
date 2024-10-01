package main

import "fmt"

func main() {
	//first()
	fmt.Println("------------------------------------")
	slices()
}
func first(){
	var buffer [256]byte
	fmt.Println(buffer[0])
	//fmt.Println(buffer[257])
	fmt.Printf("%T, %v", buffer, buffer)
	fmt.Println(len(buffer))
	var slice []byte = buffer[100:150]
	/*
		type sliceHeader struct {
		Length        int
		ZerothElement *byte
		}

		slice := sliceHeader{
			Length:        50,
			ZerothElement: &buffer[100],
		}
	*/
	//var slice2 []byte = slice[5:10]
	/*
		slice2 := sliceHeader{
		Length:        5,
		ZerothElement: &buffer[105],
		}
	*/
	slice = slice[5:10]
	var str string = "hello"
	for i := 0; i < len(str); i++ {
		fmt.Printf("%d, %c\n", i, str[i])
	}
	var i int = 0
	for i < len(str) {
		fmt.Printf("%d, %c\n", i, str[i])
		i++
	}
	i = 0
	for {
		fmt.Printf("%d, %d\n", i, str[i])
		i++
		if i >= len(str) {
			break
		}
	}
	for i, v := range str {
		fmt.Printf("%d, %c\n", i, v)
			/*
			if a := len(str) - i - 2; a == 0{
				break
			} else{
				fmt.Println(a)
			}
				*/
		
	}
}
func slices() {
	var buffer [10]int
	slice := buffer[0:0]
	for i := 0; i < len(buffer); i++ {
		n := len(slice)
		slice = slice[0 : n+1]
		slice[n] = i
		fmt.Println(slice)
	}

	/*
			type sliceHeader struct {
		    Length        int
		    Capacity      int
		    ZerothElement *byte
			}
			slice := sliceHeader{
		    Length:        0,
		    Capacity:      10,
		    ZerothElement: &iBuffer[0],
			}
	*/
	fmt.Printf("%d, %d\n", len(slice), cap(slice))
	//slice = make([]int, len(buffer), 2 * len(buffer))//type, length, capacity
	slice2 := make([]int, 15)
	copy(slice2, slice)
	slice = slice2
	slice2 = append(slice2, 5)
	slice2 = append(slice2, 3)
	slice2 = append(slice2, 2)
	slice2 = append(slice2, 7, 9, 10, 12, 15)
	slice2 = append(slice2, slice...)
	for i, v := range slice2 {
		fmt.Printf("%d, %d\n", i, v)
	}
	slice3 := []int{1, 2, 3, 4, 5}
	fmt.Printf("%T, %v", slice3, slice3)
}
