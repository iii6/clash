package main

import (
	"fmt"
	"os"
)

func main() {
	nullFile("./v2ray.txt", []byte(""))
	nullFile("./clash.yml", []byte(""))

	fmt.Println("ok.")
}

func nullFile(f string, body []byte) {
	if err := os.Remove(f); err != nil {
		panic(err)
	}

	file, err := os.Create(f)
	if err != nil {
		panic(err)
	}

	if _, err := file.Write(body); err != nil {
		panic(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)
}
