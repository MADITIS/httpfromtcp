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

	// current_line :=
	buf := make([]byte, 8)

	// read the file (8 bytes at a time) into the bytes slice
	// get the line until new lines happens
	// print that line
	var line []byte

	for {
		content, err := file.Read(buf)
		// if (io.EOF == err) {
		// 	break
		// }
		if (err != nil) {
			break
		}
		for i := 0; i < content; i++ {
			if (buf[i] == '\n') {
				fmt.Printf("read: %s\n", string(line))
				line = []byte{}
			} else {
				line = append(line, buf[i])
			}
		}




	}
	if len(line) > 0 {
        fmt.Printf("read: %s\n", string(line))
    }

}