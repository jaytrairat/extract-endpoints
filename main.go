package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	searchStrs := []string{"ws://", "wss://"}
	filepath.Walk("./CodeExtract", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("Folder not found")
		} else {
			if !info.IsDir() && (filepath.Ext(path) == ".java" || filepath.Ext(path) == ".kt" || filepath.Ext(path) == ".xml" || filepath.Ext(path) == ".html") {
				file, _ := os.Open(path)
				scanner := bufio.NewScanner(file)

				for scanner.Scan() {
					line := scanner.Text()
					for _, searchStr := range searchStrs {
						if strings.Contains(line, searchStr) {
							fmt.Printf("%s\n", strings.TrimSpace(line))
						}
					}
				}

			}
		}
		return nil
	})
}
