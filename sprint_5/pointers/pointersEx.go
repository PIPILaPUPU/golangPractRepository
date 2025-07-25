package pointers

import "fmt"

// Reverse переставляет символы строки в обратном порядке.
func Reverse(s string) string {
	if len(s) == 0 {
		return ""
	}
	// преобразуем строку в слайс рун
	runes := []rune(s)
	end := len(runes)
	for i := 0; i < end/2; i++ {
		// движемся одновременно с начала и с конца
		// и меняем местами значения элементов
		runes[i], runes[end-i-1] = runes[end-i-1], runes[i]
	}
	// приводим слайс рун к строке
	return string(runes)
}

func exequte() {
	s := "привет"
	s = Reverse(s)
	fmt.Println(s)
}
