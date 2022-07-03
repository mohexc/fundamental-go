package gobasic

import (
	"errors"
	"math"
)

func sqrt(value float64) (float64, error) {
	if value < 0 {
		return 0, errors.New("sqrt of negative number")
	}
	return math.Sqrt(value), nil
}

func ShowErrorHandler() {
	result, err := sqrt(-1)
	if err != nil {
		println(err.Error())
		return
	}
	println(result)
}
