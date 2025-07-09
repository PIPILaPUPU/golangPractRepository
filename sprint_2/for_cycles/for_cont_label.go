package forcycles

import "fmt"

func for_cycles() {
top:
	for lenght := 100; lenght <= 300; lenght++ {
		for height := 100; height <= 300; height++ {

			if V := lenght * lenght * height; V == 5_000_000 {
				fmt.Println(lenght, lenght, height, V)
				break top
			}

		}
	}
}
