package main

import (
	"os"
	"path/filepath"
)

func main() {
	_ = filepath.Walk(".", walkFn)
}

func walkFn(path string, info os.FileInfo, err error) (e error) {
	println(info.Name())
	return err
}
