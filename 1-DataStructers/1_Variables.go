package DataStructers

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/rodaine/table"
	"reflect"
	"unsafe"
)

func variables() {
	const z = 1
	var a int8 = 1
	var b rune = 'a'
	var c bool
	var d float64

	var (
		e = func() {
			fmt.Println("Coffee Go 1")
		}
		f = func() int {
			fmt.Println("Coffee Go 2")
			return 2
		}()
	)

	var x int32 = 34
	var r int16 = 12

	_ = int32(r) + x

	g := "Coffee Go"

	tbl := table.New("Type", "Value", "Address", "Size")
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	tbl.AddRow(reflect.TypeOf(a), a, &a, unsafe.Sizeof(a))
	tbl.AddRow(reflect.TypeOf(b), fmt.Sprintf("%c", b), &b, unsafe.Sizeof(b))
	tbl.AddRow(reflect.TypeOf(c), c, &c, unsafe.Sizeof(c))
	tbl.AddRow(reflect.TypeOf(d), d, &d, unsafe.Sizeof(d))
	tbl.AddRow(reflect.TypeOf(f), f, &f, unsafe.Sizeof(f))
	tbl.AddRow(reflect.TypeOf(e), e, &e, unsafe.Sizeof(e))
	tbl.AddRow(reflect.TypeOf(g), g, &g, unsafe.Sizeof(g))

	tbl.Print()
}

func f1() {
	b := 8
	c := f2()
	_ = *c + b
}

func f2() *int {
	var a = 12
	return &a
}
