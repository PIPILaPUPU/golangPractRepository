package sprint_3

import "fmt"

// Далее создаем функцию, которая будет принимать функцию обратного вызова
// в качестве параметра `callback`
func Process(numbers []int, callback func(int) int) []int {
	for i, num := range numbers {
		// Вызываем функцию, переданную в параметре для каждого элемента в слайсе
		numbers[i] = callback(num)
	}

	return numbers
}

func callBackFunc() {
	// Передаем в Process() слайс и анонимную функцию вместе с ее реализацией
	res := Process([]int{1, 2, 3}, func(num int) int {
		return num * num
	})

	fmt.Println(res)
}
