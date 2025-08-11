package main

import (
	"fmt"
	"net/url"
)

func FormValues(achievement string) url.Values {
	// переделайте код функции с использованием переменной типа url.Values
	// и методов Set() и Add()
	value := url.Values{}

	value.Set("name", "Вася")
	value.Set("nick", "superstar")
	value.Set("achieves", "cool")
	value.Add("achieves", "best")
	value.Add("achieves", achievement)

	return value
}

func main() {
	vals := FormValues("80 level")

	fmt.Println(vals["name"])
	fmt.Println(vals["nick"])
	fmt.Println(vals["achieves"])
}
