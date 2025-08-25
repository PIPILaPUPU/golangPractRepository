package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func Dir(path, shift string) error {
	// получаем список файлов и поддиректорий
	files, err := os.ReadDir(path)
	if err != nil {
		return err
	}
	for _, file := range files {
		fi, err := file.Info()
		if err != nil {
			return err
		}
		if file.IsDir() {
			fmt.Printf("%s%s\n", shift, file.Name())
			// заходим внутрь директории
			err = Dir(filepath.Join(path, file.Name()), shift+"  ")
			if err != nil {
				return err
			}
			continue
		}
		// выводим информацию о файле
		fmt.Printf("%s%s %s %d\n", shift, file.Name(),
			fi.ModTime().Format("02.01.06 15:04:05"), fi.Size())
	}
	return nil
}

func SizeDir(path string) (int64, error) {
	var sum int64
	files, err := os.ReadDir(path)
	if err != nil {
		return 0, err
	}

	for _, file := range files {
		fi, err := file.Info()
		if err != nil {
			return 0, err
		}

		if fi.IsDir() {
			var size int64

			size, err := SizeDir(filepath.Join(path, file.Name()))
			if err != nil {
				return 0, err
			}

			sum += size
			continue
		}

		sum += fi.Size()
	}

	return sum, nil
}

func main() {
	err := os.Chdir("..")
	curDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Директория:", curDir)
	if err = Dir(curDir, ""); err != nil {
		log.Fatal(err)
	}
}
