package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	filepath.Walk("./CodeExtract", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("Folder not found")
		} else {
			if !info.IsDir() && filepath.Ext(path) == ".java" {
				fmt.Println(path)
			}
		}
		return nil
	})
}
