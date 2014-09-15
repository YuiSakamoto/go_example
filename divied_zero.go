package main

import (
	"errors"
	"fmt"
	"log"
)

func div(i, j int) (result int, err error) {
	if j == 0 {
		// 自作のエラーを返す
		err = errors.New("divied by zero")
		return
	}

	result = i / j
	return
}

func main() {
	n, err := div(10, 0)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(n)
}
