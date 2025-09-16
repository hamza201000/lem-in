package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	fmt.Println(len(args))
	if len(args) == 1 {
		if strings.HasSuffix(args[0], ".txt") {
			os.Create(args[0])
		} else {
			fmt.Println("you have to write command like this : go run . exfile.txt")
			return
		}
		file, err := os.Open(args[0])
		if err != nil {
			fmt.Println("Error: ", err)
		}
		defer file.Close()
		scanneer := bufio.NewScanner(file)
		for scanneer.Scan() {
		}
	}
}
