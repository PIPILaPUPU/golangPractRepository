package main

import (
	"fmt"
	"math"
	"math/rand"
)

func Winner(rnd *rand.Rand, bets map[int]string) string {
	// напишите код функции
	var winner string
	winners := make(map[int][]string, 5)
	maxVal := 10000

	for key, val := range bets {
		diff := math.Abs(float64(key - rnd.Int()%100))

		if diff <= float64(maxVal) {
			maxVal = int(diff)
			winners[int(diff)] = append(winners[int(diff)], val)
			winner = val
		}
	}
	return winner
}

func main() {
	rnd := rand.New(rand.NewSource(1001))
	// кто какое число загадал
	bets := map[int]string{
		20: "Маша",
		34: "Игорь",
		77: "Олег",
		51: "Света",
		2:  "Саша",
	}

	fmt.Println("Победитель:", Winner(rnd, bets))
}
