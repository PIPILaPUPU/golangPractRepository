package sprint_3

import "fmt"

// объявите функцию roomsEqual() с параметрами roomSize и roomList
func roomsEqual(roomSize float64, roomList []float64) {
	// перенесите следующий код в тело функции, которую вы объявили
	count := 0
	for _, room := range roomList {
		if room == roomSize {
			count = count + 1
		}
	}
	fmt.Printf("Комнат площадью %.2f кв. м: %d\n", roomSize, count)
}

func sprint_3_1() {
	flat := []float64{
		5.55, 22.19, 7.78, 26.86, 5.55,
		29.84, 22.19, 5.55, 16.85, 4.52,
	}
	hut := []float64{9.2, 3.5, 8.1, 2.3, 9.2, 4.2, 6.9}

	roomsEqual(5.55, flat)
	// добавьте ещё один вызов функции roomsEqual()
	// передайте в функцию искомую площадь 9.2 и слайс hut
	roomsEqual(9.2, hut)
}

/*
также можно вызываьт функцию следующим образом

func func_name(words ...string)

эта запись получает неопределенное кол-во переменных
и с ним можно работать как со слайсом, например с помощью range

for _, word := range words {

}
*/
