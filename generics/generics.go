package main

import (
	"fmt"
	"math"

	"golang.org/x/exp/constraints"
)

type Tree[T any] struct {
	parent *Tree[T]
	left   *Tree[T]
	right  *Tree[T]
	val    T
}

type Number interface {
	uint | float64 | ~int
}
type MyNum int

func main() {
	a := []int{1, 2, 3, 4, 5}
	b := []int{5, 4, 3, 2, 1}
	fmt.Println(swap(&a, &b))
	fmt.Println(a, b)
	fmt.Println(find(a, 2))
	var root Tree[int] = Tree[int]{nil, nil, nil, 1}
	root.left = &Tree[int]{&root, nil, nil, 2}
	root.right = &Tree[int]{&root, nil, nil, 3}
	fmt.Println(root, *root.left, *root.right)
	m := map[int]float64{0: 1, 1: 2.3, 2: -3.3}
	n := map[int]float64{0: 2, 1: 2, 2: 2}

	//fmt.Println(addMaps[int, float64](m, n))
	fmt.Println(addMaps(m, n)) //type inference
	c := map[int]MyNum{0: 1, 1: 2, 2: 3}
	d := map[int]MyNum{0: 3, 1: 2, 2: 1}
	fmt.Println(addMaps(c, d))
}
func swap[T any](a *T, b *T) (c *T, d *T) {
	c = b
	d = a
	return
	//naked return
}
func find[T comparable](slice []T, index T) int {
	for i, v := range slice {
		if v == index {
			return i
		}
	}
	return -1
}

func addMaps[K Number, V Number](m map[K]V, n map[K]V) map[K]V {
	counter, i, j := K(0), K(0), K(0)
	var newMap map[K]V = make(map[K]V, int(math.Max(float64(len(m)), float64(len(n)))))
	for {
		if int(i) >= len(m) || int(j) > len(n) {
			break
		}
		newMap[counter] = m[i] + n[j]
		i++
		j++
		counter++
	}
	return newMap
}

func min[T constraints.Ordered](x, y T) T {
	if x < y {
		return x
	}
	return y
}
