package sprint_3

import (
	"fmt"
	"slices"
	"sort"
)

// Mode возвращает моды числовой последовательности.
// Напишите код функции
func Mode(sl []int) ([]int, int) {
	sort.Ints(sl)
	hashTable := make(map[int]int, len(sl))

	mapVal := make([]int, 0, len(sl))
	mapKey := make([]int, 0, len(sl))
	retSlice := make([]int, 0, len(sl))

	if len(sl) != 0 {

		for _, val := range sl {
			_, ok := hashTable[val]

			if !ok {
				hashTable[val] = 1
				continue
			}

			hashTable[val]++
		}

		for key, val := range hashTable {
			mapKey = append(mapKey, key)
			mapVal = append(mapVal, val)
		}

		if slices.Max(mapVal) != 1 {
			for key := range hashTable {
				if hashTable[key] == slices.Max(mapVal) {
					retSlice = append(retSlice, key)
				}
			}

			return retSlice, slices.Max(mapVal)
		}

		return []int{}, slices.Max(mapVal)

	}

	return []int{}, 1
}

func sprint3() {
	lists := [][]int{
		{},
		{57},
		{78, -7},
		{99, 200, 0},
		{4, 4, 4, 3},
		{102, -7, 44, -7, 102},
		{82, -23, 1, 5, 98, 100},
		{100000, 90000, 20000, 20000, 20000, 22000, 25500, 22000},
	}
	modes := [][]int{
		{}, {}, {}, {},
		{4},
		{-7, 102}, {},
		{20000},
	}
	mcount := []int{
		1, 1, 1, 1, 3, 2, 1, 3,
	}
	for i, list := range lists {
		mode, count := Mode(list)
		if len(mode) != len(modes[i]) {
			fmt.Printf("len mode %d: %v != %v'\n", i, modes[i], mode)
		} else {
			for j, v := range mode {
				if v != modes[i][j] {
					fmt.Printf("mode %d: %v != %v\n", i, modes[i], mode)
				}
			}
		}
		if count != mcount[i] {
			fmt.Printf("mcount %d: %d != %d\n", i, mcount[i], count)
		}
	}
	// for _, list := range lists {
	// 	l1, count := Mode(list)
	// 	fmt.Println(l1, count)
	// }
	fmt.Println("Тестирование завершено")
}
