package gobasic

import "fmt"

func ShowMap() {
	// create map
	var map1 map[string]int
	fmt.Println("create map1:", map1)

	// define map
	map1 = map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	fmt.Println("define map1:", map1)

	// length of map
	fmt.Println("lenght of map1:", len(map1))

	// update value of map
	map1["a"] = 100
	fmt.Println("update map1[a]:", map1)

	// delete value of map
	delete(map1, "a")
	fmt.Println("delete map1[a]:", map1)

	// delete map
	map1 = map[string]int{}
	fmt.Println("delete map1:", map1)

	// add element to map
	map1["a"] = 1
	map1["b"] = 2
	map1["c"] = 3
	fmt.Println("add element to map1:", map1)

	// loop map
	for k, v := range map1 {
		fmt.Println("map1[", k, "]:", v)
	}
	// foreach map
	for k, v := range map1 {
		fmt.Println("map1[", k, "]:", v)
	}

}
