package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/nwaples/rardecode"
)

func main() {
	// walk dir.
	start := time.Now()
	log.Println("search for capacitor fail. Result in d:/agilent_ICT")
	err := filepath.Walk(".", walkfunc)
	if err != nil {
		panic("walk error")
	}
	elaspe := time.Since(start)
	log.Printf("Total time usage %s", elaspe)
}

func walkfunc(path string, info os.FileInfo, err error) (e error) {
	// search rar
	if strings.HasPrefix(path, "sff2-47") && strings.HasSuffix(strings.ToLower(path), ".rar") {
		log.Printf("processing %s", path)
		r, err := rardecode.OpenReader(path, "")
		if err != nil {
			log.Printf("Error open %s", path)
		}
		defer r.Close()
		defer log.Println("")

		for {
			f, err := r.Next()
			if err != nil {
				break
			}

			content, err := ioutil.ReadAll(r)
			if err != nil {
				log.Printf("error read file %v", content)
			}

			// file c2 or c3 fail
			cont := string(content)
			if strings.Contains(cont, "%c2|01") || strings.Contains(cont, "%c3|01") {
				if strings.Contains(cont, "%pins|0") {
					log.Printf("%s", f.Name)

					// if found copy to another dir
					fo := filepath.Join("d:/", f.Name)
					os.MkdirAll(filepath.Dir(fo), 0777)
					ioutil.WriteFile(fo, content, 0777)
					if err != nil {
						log.Printf("cannot crate %s, %s", fo, err)
					}
				}
			}

		}

	}
	return nil
}
