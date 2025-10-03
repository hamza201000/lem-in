package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	start      = []string{}
	end        = []string{}
	Room_Start = ""
	Room_End   = ""
	the_rooms  = make(map[string][]string)
	visit      = make(map[string]bool)
	RestOfAnt  = 0
	tunnels    = [][]string{}

	// len_path = make(map[int]int)
)

func Ant_Path(start, end string) {
	path := []string{}
	Find_Path(start, end, path)
}

func Find_Short_Path() {
	for i := 0; i < len(tunnels); i++ {
		for j := 0; j < len(tunnels); j++ {
			if len(tunnels[i]) < len(tunnels[j]) {
				tunnels[i], tunnels[j] = tunnels[j], tunnels[i]
			}
		}
	}
}

// func Enter_Ant(){

// 	for _,path:=range tunnels{
// 		for i:=0;i<len(path);i++{

// 		}

// 	}

// }

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
	// fmt.Printf("Backtracking from %s, Path before backtrack: %v\n", current, path)
}

func Split_Char(str []byte, char byte) int {
	for i, c := range str {
		if char == c {
			return i
		}
	}
	return -1
}

func Relation_Room(Firstroom, neighbors []byte) {
	if the_rooms[string(Firstroom)] == nil {
		neighbr := []string{string(neighbors)}
		the_rooms[string(Firstroom)] = neighbr
	} else {
		the_rooms[string(Firstroom)] = append(the_rooms[string(Firstroom)], string(neighbors))
	}
	if the_rooms[string(neighbors)] == nil {
		neighbr := []string{string(Firstroom)}
		the_rooms[string(neighbors)] = neighbr
	} else {
		the_rooms[string(neighbors)] = append(the_rooms[string(neighbors)], string(Firstroom))
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

		number_of_ants := -1
		OnStart := false
		OnEnd := false
		for scanner.Scan() {
			text := scanner.Text()
			if number_of_ants == -1 {
				number_of_ants, err = strconv.Atoi(text)
				if err != nil {
					fmt.Println("Error: ", err)
					return
				}
				fmt.Println(text)
			} else if text == "##start" {
				start = append(start, text)
				OnStart = true
			} else if text == "##end" {
				end = append(end, text)
				OnEnd = true
				OnStart = false
			} else if OnStart {
				start = append(start, text)
			} else if OnEnd {
				end = append(end, text)
			}
		}
		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading file:", err)
			return
		}
		// for key, strr := range the_rooms {
		// 	fmt.Printf("%s-%s ", key, strr)
		// }
		// Ant_Path("1", "0")
		// fmt.Println()
		i := 0
		for _, ValueStart := range start {
			fmt.Println(ValueStart)
		}
		fmt.Println("endendendendendendendendendendendendendendendendendendendendendend")
		for _, ValueEnd := range end {
			fmt.Println(ValueEnd)
			ValBytes := []byte(ValueEnd)
			i := bytes.IndexByte(ValBytes, '-')
			if i != -1 {
				Relation_Room(ValBytes[:i], ValBytes[i+1:])
			}
		}
		i = Split_Char([]byte(start[1]), ' ')
		Room_Start = start[1][:i]
		i = Split_Char([]byte(end[1]), ' ')
		Room_End = end[1][:i]
		fmt.Println(Room_Start, ",", Room_End)

		Ant_Path(Room_Start, Room_End)

		fmt.Println("ok")
		fmt.Println(len(start))
		fmt.Println(len(end))
		// fmt.Println(len(tunnels))
		Find_Short_Path()
		for i, path := range tunnels {
			fmt.Println(i, path)
		}

	}
}
