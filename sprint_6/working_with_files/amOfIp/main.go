package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Config struct {
	Time string `json:"time"`
	Ip   string `json:"ip"`
	Url  string `json:"url"`
}

var cfg = Config{
	Time: "00.00.00 00:00:00",
	Ip:   "0.0.0.0",
	Url:  "",
}

func main() {
	data := make(map[string]int, 6)

	file, err := os.OpenFile("access.log", os.O_RDONLY, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	resFile, err := os.OpenFile("result.txt", os.O_RDWR, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer resFile.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		err = json.Unmarshal([]byte(scanner.Text()), &cfg)
		if err != nil {
			log.Fatal(err)
		}

		if _, ok := data[cfg.Ip]; !ok {
			data[cfg.Ip] = 1
			continue
		}

		data[cfg.Ip] += 1
	}

	stdout := os.Stdout

	os.Stdout = resFile

	for key, val := range data {
		fmt.Println(key, val)
	}

	os.Stdout = stdout
}
