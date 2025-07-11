package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func Counter(count int, t time.Time) string {
	date := t.Format("02.01.2006")
	return fmt.Sprintf("%d %s", count, date)
}

func ParseCounter(input string) (int, time.Time, error) {
	// допишите код функции
	out := strings.Split(input, " ")
	count, err := strconv.Atoi(out[0])
	data, err := time.Parse("02.01.2006", out[1])
	return count, data, err
}

func main() {
	now := time.Now()
	for count := 1; count < 4; count++ {
		s := Counter(count, now)
		counter, t, err := ParseCounter(s)
		if err != nil {
			fmt.Println("Ошибка:", err)
			return
		}
		fmt.Println(counter, t.Format("02.01.2006"))
	}
}
