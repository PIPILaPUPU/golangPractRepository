package slices_arrays_maps

import "fmt"

func arrays() {
	arr := [5]int{11, 12, 13, 14, 15}
	s := arr[:]
	fmt.Println(`s =`, s)

	fmt.Println(`s[:2] =`, s[:2])
	fmt.Println(`s[len(s)-2:] =`, s[len(s)-2:])
	fmt.Println(`s[1:4] =`, s[1:4])

	s = s[:0] // удалили все элементы
	fmt.Println(`s[:0] =`, s)
	fmt.Println(`len(s) =`, len(s), `cap(s) =`, cap(s))
}
