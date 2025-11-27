package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("./messages.txt")

	if err != nil {
		fmt.Println("Could not read the file")

	}

	buf := make([]byte, 8)

	for {
		content, err := file.Read(buf)
		if (err != nil) {
			break
		}

		fmt.Printf("read: %s\n", string(buf[:content]))

	}

}