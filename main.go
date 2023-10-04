package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "extract-endpoints",
	Short: "Search for ws and wss in string",
	Run:   runSearch,
}

func runSearch(cmd *cobra.Command, args []string) {
	searchStrs := []string{"ws://", "wss://"}
	file, err := os.Create("result.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	filepath.Walk("./CodeExtract", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("Folder not found")
		} else {
			if !info.IsDir() && (filepath.Ext(path) == ".java" || filepath.Ext(path) == ".kt" || filepath.Ext(path) == ".xml" || filepath.Ext(path) == ".html") {
				f, err := os.Open(path)
				if err != nil {
					fmt.Println("Error:", err)
					return nil
				}
				defer f.Close()

				scanner := bufio.NewScanner(f)

				for scanner.Scan() {
					line := scanner.Text()
					for _, searchStr := range searchStrs {
						if strings.Contains(line, searchStr) {
							file.WriteString(fmt.Sprintf("%s\t%s\n", path, strings.TrimSpace(line)))
						}
					}
				}
			}
		}
		return nil
	})

	fmt.Printf("%s :: file created\n", time.Now().Format("2006-01-02 15:04:05"))
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
