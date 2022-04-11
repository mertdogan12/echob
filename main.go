package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
)

func main() {
	br := bufio.NewReader(os.Stdin)
	input := make([]byte, 0)

	for {
		line, _, _ := br.ReadLine()
		if len(line) != 0 {
			input = append(input, []byte(line)...)

			clear()
			fmt.Println(hex.Dump(input))
		}
	}
}

func clear() {
	fmt.Print("\033[H\033[2J")
}
