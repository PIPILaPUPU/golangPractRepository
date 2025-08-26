package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rnd := rand.NewSource(time.Now().Unix())
	data := make([]string, 0, 10)

	for i := 0; i < 10; i++ {
		data = append(data, strGen(10, rnd))
	}

	fmt.Print(data)
}

func strGen(n int, gen rand.Source) string {
	res := make([]byte, 0, n)

	for i := 0; i < n; i++ {
		letter := gen.Int63()%26 + 97

		res = append(res, byte(letter))
	}

	return string(res)
}
