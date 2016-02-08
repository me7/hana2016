package main

import (
	"os"
)

func main() {
	_, err := os.Open("somefile")
	if err != nil {
		println(err.Error())
	}
}
