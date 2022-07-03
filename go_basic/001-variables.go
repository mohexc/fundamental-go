package gobasic

import "fmt"

func ShowVariables() {
	var a int = 10
	var b string = "Hello"
	var c bool = true
	var d float64 = 10.5
	var e rune = 'a'
	var f complex64 = 10 + 10i
	var g interface{} = "Hello"
	var h map[string]int = map[string]int{"a": 1, "b": 2}
	var i []int = []int{1, 2, 3}
	var j func() = func() {
		println("Hello")
	}
	var k struct {
		a int
		b string
	}
	k.a = 10
	k.b = "Hello"

	fmt.Println(a, b, c, d, e, f, g, h, i, k)
	j()
}
