package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"-"`
	active   bool
}

func main() {
	user := User{
		ID:       1,
		Name:     "Гофер",
		Email:    "gopher@gophermate.com",
		Password: "I4mG0ph3R",
		active:   true,
	}

	// преобразуйте user в JSON формат
	u, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
		return
	}
	var newUser User

	// десериализуйте данные из JSON формата в переменную newUser
	err = json.Unmarshal(u, &newUser)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(newUser)
}
