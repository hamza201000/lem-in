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
	visit          = make(map[string]bool)
	RestOfAnt      = 0
	tunnels        = [][]string{}
	Enter_Ant      = []string{}
)

func Ant_Path(start, end string) {
	path := []string{}
	Find_Path(start, end, path)
}

func Find_Path(current, end string, path []string) {
	path = append(path, current)
	visit[current] = true
	if end == current {
		pathCopy := make([]string, len(path))
		copy(pathCopy, path)
		tunnels = append(tunnels, pathCopy)
		visit[current] = false
		return
	}
	for _, room := range the_rooms[current] {
		if !visit[room] {
			Find_Path(room, end, path)
		}
	}
	visit[current] = false
	fmt.Printf("Backtracking from %s, Path before backtrack: %v\n", current, path)
}

func Check_Char(str string, char rune) int {
	for i, c := range str {
		if char == c {
			return i
		}
	}
	return -1
}

func Check_slice(slc []string, str string) bool {
	for _, check := range slc {
		if str == check {
			return true
		}
	}
	return false
}

func Relation_Room(Firstroom, neighbors string) {
	RoomExist := false
	for room, neighbor := range the_rooms {
		if room == Firstroom {
			RoomExist = true
			neighbor = append(neighbor, neighbors)
			the_rooms[Firstroom] = neighbor
		}
	}
	if !RoomExist {
		var neighbor []string
		neighbor = append(neighbor, neighbors)
		the_rooms[Firstroom] = neighbor
	}
	for room, neighbor := range the_rooms {
		for _, inside := range neighbor {
			if !Check_slice(the_rooms[inside], room) {
				the_rooms[inside] = append(the_rooms[inside], room)
			}
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
					Relation_Room(string(text[:i]), string(text[i+1:]))
				}
				fmt.Println(text)
			}
		}
		for key, strr := range the_rooms {
			fmt.Printf("%s-%s ", key, strr)
		}
		Ant_Path("start", "end")
		fmt.Println()
		for _, paths := range tunnels {
			fmt.Print(paths)
			fmt.Println()
		}

		fmt.Println(len(tunnels))
		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading file:", err)
			return
		}
	}
}
