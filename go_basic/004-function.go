package gobasic

func ShowFunciton() {
	// function
	baseFunc := func() {
		println("Hello")
	}
	baseFunc()
	// function with parameter
	funcWithOneParam := func(name string) {
		println("Hello", name)
	}
	funcWithOneParam("World")
	// function with return value
	funcWithReturn := func() string {
		return "Hello"
	}
	println(funcWithReturn())
	// function with return value and parameter
	funcWithReturnAndParam := func(name string) string {
		return "Hello " + name
	}
	println(funcWithReturnAndParam("World"))
	// function with parameter and multiple return value
	funcWithParamAndMultipleReturn := func(name string) (string, string) {
		return "Hello", name
	}
	println(funcWithParamAndMultipleReturn("World"))
	// function with multiple return value and multiple parameter
	funcWithMultipleReturnAndMultipleParam := func(name string, age int) (string, int) {
		return "Hello", age
	}
	println(funcWithMultipleReturnAndMultipleParam("World", 20))
}

// function
func baseFunc() {
	println("Hello")
}

// function with parameter
func funcWithOneParam(name string) {
	println("Hello", name)
}

// function with return value
func funcWithReturn() string {
	return "Hello"
}
