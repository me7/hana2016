package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	filepath.Walk(".", walkfunc)
}

func walkfunc(path string, info os.FileInfo, err error) error {
	if info.Name() == "testplan" {
		filename, err := filepath.Abs(path)
		if err != nil {
			return err
		}
		processFile(filename)
	}
	return nil
}

func processFile(filename string) error {
	var output []string

	text, err := ioutil.ReadFile(filename)
	if err != nil {
		println(err.Error())
		return err
	}

	lines := strings.Split(string(text), "\n")

	for i := range lines {
		output = append(output, lines[i])
		if strings.HasPrefix(lines[i], "Start_Time = msec") ||
			strings.HasPrefix(lines[i], "Starttime = msec") ||
			strings.HasPrefix(lines[i], "Start_time=msec") {
			output = append(output, "c:/Agilent_ICT/ICT.exe start generic;append !gampolt")
		}
		if strings.HasPrefix(lines[i], "End_Time = msec") ||
			strings.HasPrefix(lines[i], "Endtime = msec") ||
			strings.HasPrefix(lines[i], "End_time=msec") {
			output = append(output, "c:/Agilent_ICT/ICT.exe stop generic;append !gampolt")
		}
	}

	ioutil.WriteFile(filename+".out", []byte(strings.Join(output, "\n")), 0644)

	return nil
}
