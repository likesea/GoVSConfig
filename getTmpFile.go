package main

import (
	"os"
	"path/filepath"
)

func GetTmpFile(absPath string, extension string) []string {
	var filesSlice []string

	filepath.Walk(absPath, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == extension {
			filesSlice = append(filesSlice, path)
		}
		return nil
	})

	return filesSlice
}

/*	for _, v := range filesSlice {
	if v != nil {
		//open the file
		fi, _ := os.Open(v.Name())
		//read the file
		fd, _ := ioutil.ReadAll(fi)
		fmt.Println(fi.Name())
		content := string(fd)
		for _, v := range content {
			if string(v) == "$" {
				fmt.Println(string(v))
			}

		}
	}
}*/
