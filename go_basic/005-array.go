package gobasic

import "fmt"

func ShowArrayAndSlice() {
	// create an array
	var arr [4]int
	fmt.Println("create arr:", arr)

	// define an array with initial value
	arr2 := [4]int{1, 2, 3, 4}
	fmt.Println("define arr2:", arr2)

	// lenght of array
	fmt.Println("lenght of arr:", len(arr))

	// capacity of array
	fmt.Println("capacity of arr:", cap(arr))

	// update value of array
	arr[0] = 100
	fmt.Println("update arr[0]:", arr)

	// delete value of array
	arr[0] = 0

	// delete array
	arr = [4]int{}
	fmt.Println("delete arr:", arr)

	// loop array
	for i := 0; i < len(arr); i++ {
		fmt.Println("arr[", i, "]:", arr[i])
	}
	// foreach array
	for i, v := range arr {
		fmt.Println("arr[", i, "]:", v)
	}

	fmt.Println("---------------------- slice -------------------------------")
	// create slice
	var slice []int
	fmt.Println("create slice:", slice)

	// make slice
	slice = make([]int, 5, 10)

	// define slice
	slice = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("define slice:", slice)

	// length of slice
	fmt.Println("length of slice:", len(slice))

	// capacity of slice
	fmt.Println("capacity of slice:", cap(slice))

	// update element of slice
	slice[0] = 100
	fmt.Println("update element of slice:", slice)

	// slice first eight elements of slice
	slice = slice[:8]
	fmt.Println("slice first eight elements of slice:", slice)

	// slice last six elements of slice
	slice = slice[2:]
	fmt.Println("slice last six elements of slice:", slice)

	// slice between two elements of slice
	slice = slice[2:5]
	fmt.Println("slice between two elements of slice:", slice)

	// add element to front of slice
	slice = append(slice, 0)
	fmt.Println("add element to front of slice:", slice)

	// add element to back of slice
	slice = append(slice, 11)
	fmt.Println("add element to back of slice:", slice)

}
