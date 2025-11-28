package main

import (
	"fmt"
	"io"
	"os"
)

func getLinesChannel(file io.ReadCloser) <-chan string {
	ch := make(chan string)

	go func() {
		defer close(ch)
		defer file.Close()
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
			if err != nil {
				break
			}
			for i := 0; i < content; i++ {
				if buf[i] == '\n' {
					ch <- string(line)
					line = []byte{}
				} else {
					line = append(line, buf[i])
				}
			}

		}
		if len(line) > 0 {
			ch <- string(line)
		}
	}()

	return ch
}

func main() {
	file, err := os.Open("./messages.txt")

	if err != nil {
		fmt.Println("Could not read the file")

	}

	current_line := getLinesChannel(file)

	for line := range current_line {
		fmt.Printf("read: %s\n", line)
	}

}
