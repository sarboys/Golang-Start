package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	files, _ := getFiles(".")
	fmt.Println(files)
}

func getFiles(path string) ([]string, error) {
	var files []string
	dirFiles, err := ioutil.ReadDir(fmt.Sprintf("%s", path))
	if err != nil {
		return nil, err
	}
	for _, val := range dirFiles {
		files = append(files, val.Name())
	}
	return files, nil
}
