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

func slices() {
	var a = []int16{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var b = a[:]
	var c = a[2:]
	var d = a[:4]
	var e = a[2:4]
	var f = a[2:4:6]

	fmt.Printf("a		%d		len:%d,		cap: %d\n", a, len(a), cap(a))
	fmt.Printf("a[:]		%d		len:%d,		cap: %d\n", b, len(b), cap(b))
	fmt.Printf("a[2:]		%d		len:%d,		cap: %d\n", c, len(c), cap(c))
	fmt.Printf("a[:4]		%d			len:%d,		cap: %d\n", d, len(d), cap(d))
	fmt.Printf("a[2:4]		%d				len:%d,		cap: %d\n", e, len(e), cap(e))
	fmt.Printf("a[2:4:6]	%d				len:%d,		cap: %d\n", f, len(f), cap(f))
}

func valueAndPointerSemantic() {
	//a := []int16{1, 2, 3, 4, 5, 6, 7, 8, 9}
	//for _, v := range a {
	//	a = a[:3]
	//	fmt.Println(v)
	//}
	//fmt.Println(a)

	a := []int16{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := range a {
		a = a[:3]
		fmt.Println(a[i])
	}
	fmt.Println(a)
}

func strings() {
	s := "Coffee Go"
	fmt.Println(s)
	fmt.Println(fmt.Sprintf("%c", s[1]))
	fmt.Println(s[:2])
	fmt.Println(s[3:])
	fmt.Println(s[1:4])
}

func maps() {
	m := make(map[string]int)
	m["Coffee"] = 1
	m["Go"] = 2

	fmt.Println(m, len(m))
	ElemSize(m)
}

func ElemSize(container interface{}) {
	t := reflect.TypeOf(container)
	fmt.Printf("Type:%s	ItemStize:%d	TotelSize:%d	\n", t, t.Elem().Size(), t.Size())
}
