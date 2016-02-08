package main

import (
	"fmt"
	"os"
)

var (
	filepath  = "c:/Valor/"
	startFile = "c:/Valor/StartFile"
	stopFile  = "c:/Valor/StopFile"
)

// Exists reports whether the named file or directory exists.
func Exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func main() {
	os.MkdirAll(filepath, 0777)
	fmt.Println(os.Args)
	if len(os.Args) == 1 {
		fmt.Println("use ICT.exe start/stop model")
		os.Exit(0)
	}
	if os.Args[1] == "start" {
		fmt.Println("StartFile")
		if Exists(stopFile) {
			os.Remove(stopFile)
		}
		if Exists(startFile) == false {
			f, _ := os.Create(startFile)
			f.WriteString(os.Args[2])
			f.Close()
		}
	} else if os.Args[1] == "stop" {
		fmt.Println("Stop")
		if Exists(startFile) {
			os.Remove(startFile)
		}
		if Exists(stopFile) == false {
			f, _ := os.Create(stopFile)
			f.WriteString(os.Args[2])
			f.Close()
		}
	}
}
