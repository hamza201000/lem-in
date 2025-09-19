package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	number_of_ants = -1
	first          = false
	the_rooms      = make(map[string][]string)
)

func Check_Char(str string, char rune) int {
	for i, c := range str {
		if char == c {
			return i
		}
	}
	return -1
}

func Relation_Room(graph map[string][]string, Firstroom, neighbor string) {
	for room := range graph {
		if room==Firstroom{
			graph[Firstroom]==append()
		}
	}
}

func main() {
	args := os.Args[1:]
	if len(args) == 1 {
		if strings.HasSuffix(args[0], ".txt") {
			if _, err := os.Stat(args[0]); os.IsNotExist(err) {
				_, err = os.Create(args[0])
				if err != nil {
					fmt.Println("Error: ", err)
					return
				}
			}
		} else {
			fmt.Println("you have to write command like this : go run . exfile.txt")
			return
		}
		file, err := os.Open(args[0])
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			text := scanner.Text()
			if number_of_ants == -1 {
				number_of_ants, err = strconv.Atoi(text)
				if err != nil {
					fmt.Println("Error: ", err)
					return
				}
				fmt.Println(text)
			} else if text == "##end" || text == "##start" {
				fmt.Println(text)
				first = true
			} else if first {
				i := Check_Char(text, '-')
				if i != -1 {
					fmt.Println("OK")
					Firstroom := string(text[:i])
					neighbor := string(text[i+1:])

					the_rooms[string(text[:i])] = string(text[i+1:])
				}
				fmt.Println(text)
			}
		}
		for key, strr := range the_rooms {
			fmt.Printf("%s-%s ", key, strr)
		}
		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading file:", err)
			return
		}
	}
}
