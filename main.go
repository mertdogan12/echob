package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
)

func main() {
	br := bufio.NewReader(os.Stdin)
	for {
		line, _, _ := br.ReadLine()
		if len(line) != 0 {
			fmt.Println(hex.Dump(line))
		}
	}
}
