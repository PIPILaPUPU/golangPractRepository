package sprint_4

import (
	"fmt"
	"time"
)

// формат дней рождений
const layout = "02.01.2006"

func Age(birthday string) (int, error) {
	bd, err := time.Parse("02.01.2006", birthday)
	if err != nil {
		return 0, err
	}

	if time.Now().Month() <= bd.Month() && time.Now().Day() < bd.Day() {
		return time.Now().Year() - bd.Year() - 1, nil
	}

	return time.Now().Year() - bd.Year(), nil
}

func time_pack() {
	for _, v := range []string{"04.01.1969", "28.07.1995",
		"16.12.2007", "07.03.1947"} {
		age, err := Age(v)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(v, ":", age)
	}
}
