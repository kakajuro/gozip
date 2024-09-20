package main

import (
	lib "kakajuro/gozip/lib"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	// Check if there are CLI args
	if len(os.Args) <= 1 {
		fmt.Printf("Incorrect number of arguments. Specify a file or directory to zip.")
		os.Exit(0)
	}

	// Check if provided file/directory exists
	res := lib.DoesExist(os.Args[1]);
	if res { 
		fmt.Println("Time to zip files.") 
	} else {
		os.Exit(1)
	}

	// Get file/directory path
	fileInfo, err := os.Stat(os.Args[1])
	if err != nil {
		message := fmt.Sprintf("Error retriving file %s", os.Args[1])
		fmt.Println(message)
		os.Exit(1)
	}

	// Generate filename for zip
	// Converts all slashes to / and gets the name of the last dir/file
	// then removes any file extensions
	toZip, err := filepath.Abs(fileInfo.Name())
	if err != nil {
		fmt.Println("Failed to get file/directory path")
		os.Exit(1)	
	}

	zipnameArr := strings.Split(strings.ReplaceAll(toZip, "\\", "/"), "/")
	zipnameWithExtension := zipnameArr[len(zipnameArr)-1]
	zipname := strings.Split(zipnameWithExtension, ".")[0]
	
	fmt.Println(zipname)

}
