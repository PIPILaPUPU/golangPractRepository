package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rnd := rand.NewSource(time.Now().Unix())
	data := make([]string, 0, 15)

	for i := 0; i < 15; i++ {
		data = append(data, stringGenerator(10, rnd))
	}

	fmt.Println(data)
}

func stringGenerator(n int, generator rand.Source) string {
	by := make([]byte, 0, n+1)

	for i := 0; i < n; i++ {
		by = append(by, byte(generator.Int63()%26+97))
	}

	return string(by)
}
