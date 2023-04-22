package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a = 12
	var b rune = 'a'
	var c bool
	var d float64
	var e string

	fmt.Printf("int: 	%d 			address: %p 		size: %d \n", a, &a, unsafe.Sizeof(a))
	fmt.Printf("rune: 	%d 			address: %p 		size: %d \n", b, &b, unsafe.Sizeof(b))
	fmt.Printf("bool: 	%t 			address: %p 		size: %d \n", c, &c, unsafe.Sizeof(c))
	fmt.Printf("string:	%s 			address: %p 		size: %d \n", e, &e, unsafe.Sizeof(e))

	var (
		g = func() {
			fmt.Println("kousha")
		}
		h = func() int {
			fmt.Println()
			fmt.Println("Hello Go")
			fmt.Println()
			return 2
		}()
	)

	k := "Coffee GO"

	fmt.Printf("float:	%f 		address: %p 		size: %d \n", d, &d, unsafe.Sizeof(d))
	fmt.Printf("func:	%p 		address: %p 		size: %d \n", g, &g, unsafe.Sizeof(g))
	fmt.Printf("func:	%d 			address: %p 		size: %d \n", h, &h, unsafe.Sizeof(h))
	fmt.Printf("string:	%s 		address: %p 		size: %d \n", k, &k, unsafe.Sizeof(k))
}
