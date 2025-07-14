package main

import "fmt"

type Character struct {
	Name   string // имя
	Health int    // здоровье
	Speed  int    // скорость
	Power  int    // сила
	Woman  bool   // true, если женский персонаж
}

func (character Character) String() string {
	s := fmt.Sprintf(`Имя: %s
Здоровье: %d
Скорость: %d
Сила: %d
`, character.Name, character.Health, character.Speed, character.Power)
	return s
}

type Magician struct {
	Character
	Magic int
}

// добавьте для Magician метод String()
func (magician Magician) String() string {
	s := fmt.Sprintf(`Имя: %s
Здоровье: %d
Скорость: %d
Сила: %d
Магия: %d
`, magician.Name, magician.Health, magician.Speed, magician.Power, magician.Magic)
	return s
}

func main() {
	merlin := Magician{
		Character: Character{
			Name:   "Merlin",
			Health: 100,
			Speed:  250,
			Power:  400,
		},
		Magic: 700,
	}
	fmt.Println(merlin)
}
