package main

import (
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v3"
)

type User struct {
	ID int `json:"id" yaml:"id"`
	// добавьте поля Name и Email
	Name  string `json:"name" yaml:"name"`
	Email string `json:"email" yaml:"email"`
}

func main() {
	// исходные данные в формате YAML
	yamlData := `
id: 2
name: Гофер
email: gopher@gophermate.com
`
	// промежуточная переменная типа User
	var user User
	// переменная для конвертации в JSON формат
	var jsonData []byte

	// добавьте десериализацию yamlData в user
	err := yaml.Unmarshal([]byte(yamlData), &user)
	if err != nil {
		fmt.Println(err)
		return
	}
	// добавьте сериализацию user в jsonData
	jsonData, err = json.MarshalIndent(user, "", "    ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(jsonData))
}
