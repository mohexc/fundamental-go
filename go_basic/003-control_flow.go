package gobasic

func ShowControlFlow() {
	//  condition
	if true {
		println("Hello")
	}
	//  condition with else
	if false {
		println("Hello")
	} else {
		println("World")
	}
	//  condition with else if
	if false {
		println("Hello")
	} else if true {
		println("World")
	} else {
		println("!")
	}
	// switch
	switch true {
	case true:
		println("Hello")
	case false:
		println("World")
	default:
		println("!")
	}
	// loop
	for i := 0; i < 10; i++ {
		println(i)
	}
	// loop with condition
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		println(i)
	}
	// loop with condition and continue
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		println(i)
	}
	// loop with range
	for i, v := range []int{1, 2, 3} {
		println(i, v)
	}
	// loop with range and continue
	for i, v := range []int{1, 2, 3} {
		if i == 1 {
			continue
		}
		println(i, v)
	}
	// loop with range and break
	for i, v := range []int{1, 2, 3} {
		if i == 1 {
			break
		}
		println(i, v)
	}
	// loop with range and break and continue
	for i, v := range []int{1, 2, 3} {
		if i == 1 {
			break
		} else if i == 2 {
			continue
		}
		println(i, v)
	}
	// loop with range and break and continue and return
	for i, v := range []int{1, 2, 3} {
		if i == 1 {
			break
		} else if i == 2 {
			continue
		} else if i == 3 {
			return
		}
		println(i, v)
	}
}
