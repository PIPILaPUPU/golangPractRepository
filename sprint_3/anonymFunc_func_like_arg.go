package sprint_3

import "fmt"

// Функция, которая принимает другую функцию в качестве аргумента
func squareAndPrint(fn func(int) int) {
	result := fn(5)
	fmt.Println(result)
}

// Функция, которая будет передаваться как аргумент
func square(x int) int {
	return x * x
}

func func_as_arg() {
	// Передаем функцию square в doSomething
	squareAndPrint(square)
}

func anonym_func() {
	// вызов анонимной функции,
	// в переменной res будет сохранено true, потому что 10 < 20
	res := func(a, b int) bool { return a < b }(10, 20)

	fmt.Println(res)
}

// printOnTrue напечатает s в случае, если функция f вернёт true,
// в противном случае она ничего не сделает
func printOnTrue(f func(a int, b int) bool, s string) {
	if f(10, 20) {
		fmt.Println(s)
	}
}

func func_as_anonym_var() {
	// создаём анонимную функцию и присваиваем переменной f
	f := func(a, b int) bool { return a < b }
	// передаём функцию f в качестве аргумента функции printOnTrue
	printOnTrue(f, "и правда меньше")
}
