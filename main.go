package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func doesExist(path string) int {

	_, err := os.Stat(path)
	if err != nil {

		if os.IsNotExist(err) {
			fmt.Println("Desired file or directory does not exist")
			os.Exit(1)
		} else {
			fmt.Println(err)
			os.Exit(1)
		}

	}

	return 0

}

func main() {

	// Check if there are CLI args
	if len(os.Args) <= 1 {
		fmt.Printf("Incorrect number of arguments. Specify a file or directory to zip.")
		os.Exit(0)
	}

	// Check if provided file/directory exists
	res := doesExist(os.Args[1]);
	if res == 0 { fmt.Println("Time to zip files.") }

	// Get file/directory path
	fileInfo, err := os.Stat(os.Args[1])
	if err != nil {
		message := fmt.Sprintf("Error retriving file %s", os.Args[1])
		fmt.Println(message)
		os.Exit(1)
	}

	toZip, err := filepath.Abs(fileInfo.Name())
	if err != nil {
		fmt.Println("Failed to get file/directory path")
		os.Exit(1)	
	}

}
