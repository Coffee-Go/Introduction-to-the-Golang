package DataStructers

import "fmt"

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
	a := []int16{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, v := range a {
		a = a[:3]
		fmt.Println(v)
	}
	fmt.Println(a)

	a = []int16{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := range a {
		a = a[:3]
		fmt.Println(a[i])
	}
	fmt.Println(a)
}
