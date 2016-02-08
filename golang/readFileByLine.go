package main

import (
	"io/ioutil"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("readFileByLine.go")
	if err != nil {
		println(err.Error())
	}

	lines := strings.Split(string(data), "\n")
	for i := range lines {
		println(lines[i])
	}
}
