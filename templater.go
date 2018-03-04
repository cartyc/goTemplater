package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {

	GetTemplates("templates")

}

func GetTemplates(path string) {

	/*
		What Do I Want here?
		Loop over dir and find folders containing templates
		Create Templates
		Folder structure
		main
		-- group
		-- *.html
		- group
		-- *.html
		*.html
	*/

	var files []string
	fmt.Println(path)

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		fmt.Println(info.IsDir())
		files = append(files, path)
		return nil
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(files)
}
