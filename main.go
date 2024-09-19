package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) <= 1 {
		fmt.Printf("Incorrect number of arguments. Specify a file or directory to zip.")
		os.Exit(0)
	}

	fmt.Printf("Time to zip files.")

}
