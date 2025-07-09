package slices_arrays_maps

import "fmt"

func maps_1() {
	// оценки по отдельным ученикам
	marks := map[string][]int{
		"Светлана":  {5, 4, 5, 5, 4},
		"Артём":     {3, 4, 4, 5, 3},
		"Александр": {2, 3, 3, 4},
		"Ольга":     {5, 5, 4, 4},
		"Мария":     {4, 3, 4, 4, 3, 5},
	}
	student := "Светлана"
	var (
		average float32 // средний балл
		counter int
	)

	// допишите недостающий код
	marks[student] = append(marks[student], 5)

	for _, value := range marks[student] {
		average += float32(value)
		counter++
	}

	average = average / float32(counter)
	fmt.Printf("%.2f\n", average)
}

func maps_2() {
	// ассортимент выпечки
	bakery := map[string]int{
		"Хлеб белый":          43,
		"Хлеб Дарницкий":      47,
		"Батон":               52,
		"Булочка Домашняя":    35,
		"Хачапури":            63,
		"Сосиска в тесте":     70,
		"Беляш":               82,
		"Растегай":            87,
		"Самса":               91,
		"Пирожок с картошкой": 37,
	}
	// список покупок
	buy := []string{"Растегай", "Самса", "Хачапури", "Хлеб Подовый", "Беляш"}
	// вывод общего количества позиций в bakery
	fmt.Println("Всего позиций:", len(bakery))

	// вставьте недостающий код
	var (
		sum          int
		count_falier int
		count        int
	)

	for _, prod := range buy {
		_, ok := bakery[prod]

		if ok {
			count++
			sum += bakery[prod]
			continue
		}

		count_falier++
	}

	// вставьте вторым параметром соответствующие переменные
	fmt.Println("Ошибок:", count_falier) // количество ошибочных товаров
	fmt.Println("В корзине:", count)     // количество корректных товаров
	fmt.Println("К оплате:", sum)
}

func maps_3() {
	titles := map[string][]string{
		"Что делать?":               {"планы", "думы"},
		"Где отдохнуть в выходные":  {"отдых", "планы"},
		"Кто виноват?":              {"поиск", "думы", "общество"},
		"Лучшие рестораны города":   {"еда", "отдых"},
		"С заботой о народе":        {"думы", "общество"},
		"Как попасть на дегустацию": {"еда", "планы", "отдых"},
	}
	tags := make(map[string][]string)
	stats := make([]string, 0, 10)

	for stat, arr := range titles {
		for _, value := range arr {
			_, ok := tags[value]

			if !ok {
				stats = append(stats, stat)
			}

			tags[value] = append(tags[value], stat)
		}
	}

	fmt.Println(len(tags["думы"]), len(tags["общество"]), tags["поиск"])
}
