package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

func main() {
	path := "C:\\Go"                   // целевая директория
	files, err := ioutil.ReadDir(path) // тип возвращаемого значения - []fs.FileInfo
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() {
			fmt.Printf("%s is directory. Path: %v\n",
				file.Name(), filepath.Join(path, file.Name()))
		} else {
			fmt.Printf("%s is file. Path: %v\n",
				file.Name(), filepath.Join(path, file.Name()))
		}
	}
}
