package sprint_3

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
)

// создаем переменную ошибки тут
var (
	ErrConvToFloat       = errors.New("ошибка преобразования строки в число с плавающей точкой")
	ErrNotPositiveNumber = errors.New("число должно быть больше нуля")
)

func div() (int, error) {
	// случайное целое число от 0 до 2
	a := rand.Intn(3)
	// если а равна 0
	if a == 0 {
		// возвращаем в первую переменную 0, а во вторую - ошибку, выходим из функции
		return 0, errors.New("делитель не должен быть равен нулю")
	}
	// если а не равна 0, в первую переменную записываем результат деления, а во вторую nil, что означает, что ошибок нет
	return 1 / a, nil
}

func Error() {
	// вызываем функцию div() и результат и ошибку присваиваем переменным result и err
	result, err := div()
	// проверяем, вернулась ли ошибка
	if err != nil {
		// если да, то выводим ее в консоль и завершаем программу
		fmt.Println(err.Error())
		return
	}
	// если ошибки нет, выводим результат деления
	fmt.Println(result)
}

// ниже ваша функция сложения
func addPositive(num1, num2 string) (float64, error) {
	result := 0.0
	fl1, err := strconv.ParseFloat(num1, 64)
	if err != nil {
		return result, ErrConvToFloat
	}

	fl2, err := strconv.ParseFloat(num2, 64)
	if err != nil {
		return result, ErrConvToFloat
	}

	if fl1 <= 0.0 || fl2 <= 0.0 {
		return result, ErrNotPositiveNumber
	}

	result = fl1 + fl2
	return result, nil
}

func test() {
	num1 := "5.2"
	num2 := "2.7"

	result, err := addPositive(num1, num2)

	if err != nil {
		err = fmt.Errorf("ошибка в ходе выполнения программы: %v", err)
		fmt.Println(err)
	} else {
		fmt.Println("результат сложения", result)
	}

	num1 = "two"
	num2 = "five"

	result, err = addPositive(num1, num2)
	if err != nil {
		err = fmt.Errorf("ошибка в ходе выполнения программы: %v", err)
		fmt.Println(err)
	} else {
		fmt.Println("результат сложения:", result)
	}

	num1 = "5.4"
	num2 = "0.0"

	result, err = addPositive(num1, num2)
	if err != nil {
		err = fmt.Errorf("ошибка в ходе выполнения программы: %v", err)
		fmt.Println(err)
	} else {
		fmt.Println("результат сложения:", result)
	}
}
