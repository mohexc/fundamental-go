package gobasic

import "fmt"

func ShowPointer() {
	// create pointer
	var ptr *int
	fmt.Println("create pointer:", ptr)

	// assign value to pointer
	ptr = new(int)
	fmt.Println("assign value to pointer:", *ptr)

	// read position ram pointer
	fmt.Println("read position ram pointer:", &ptr)

	// update value of pointer
	*ptr = 100
	fmt.Println("update value of pointer:", *ptr)

}
