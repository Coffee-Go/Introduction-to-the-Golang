package DataStructers

import "fmt"

func maps() {
	m := make(map[string]int)
	m["Coffee"] = 1
	m["Go"] = 2

	fmt.Println(m, len(m))
	ElemSize(m)
}
