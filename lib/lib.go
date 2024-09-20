package lib

import (
	"fmt"
	"os"
)

func DoesExist(path string) bool {

	_, err := os.Stat(path)
	if err != nil {

		if os.IsNotExist(err) {
			fmt.Println("Desired file or directory does not exist")
			return false
		} else {
			fmt.Println(err)
			return false
		}

	}

	return true

}