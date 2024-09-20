package lib

import (
	"fmt"
	"os"
)

func DoesExist(path string) int {

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