// testcode project main.go
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func getAllDirs(oldpath string, dirs *[]string, newdirs *[]string, newpath string) {

	err := filepath.Walk(oldpath, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			*dirs = append(*dirs, path)
			*newdirs = append(*newdirs, strings.Replace(path, oldpath, newpath, 1))
		}
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}

}

func main() {
	//fmt.Println("test")

	//testpath := "D:\\temp"
	testpath := os.Args[1]
	newpath := os.Args[2]
	dirs := []string{}
	newdirs := []string{}
	getAllDirs(testpath, &dirs, &newdirs, newpath)

	for _, value := range newdirs {
		path, _ := filepath.Abs(value)
		if err := os.Chdir(path); err == nil {
			continue
		}
		os.Mkdir(path, os.ModePerm)
	}
}
