package main

import (
	"fmt"
	"errors"
)

var divByZero = errors.New("нельзя делить на ноль")

func div(x, y int) (int, error) {
	if y == 0 {
		return 0, divByZero
	}

	return x/y, nil
}

func main() {
	var a,b int
	fmt.Println(div(a, b))
	a, b = 27,3
	fmt.Println(div(a, b))
}
