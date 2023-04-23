package DataStructers

import (
	"fmt"
	"reflect"
)

func arrays() {
	var a [2]int8

	a[0] = 1
	a[1] = 2

	ElemSize(a)
	fmt.Println(len(a), cap(a))
	fmt.Println(a)
	fmt.Println()

	b := []int8{1, 2, 3, 4, 5}
	ElemSize(b)
	fmt.Println(len(b), cap(b))
	fmt.Println(b)
	fmt.Println()

	b = append(b, 6)
	ElemSize(b)
	fmt.Println(len(b), cap(b))
	fmt.Println(b)
	fmt.Println()

	c := make([]int8, 3, 6)
	copy(c, b)
	c = append(c, 8)
	ElemSize(c)
	fmt.Println(len(c), cap(c))
	fmt.Println(c)
	fmt.Println()
}

func ElemSize(container interface{}) {
	t := reflect.TypeOf(container)
	fmt.Printf("Type:%s	ItemStize:%d	TotelSize:%d	\n", t, t.Elem().Size(), t.Size())
}
