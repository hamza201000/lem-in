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
	start          = []string{}
	end            = []string{}
	Room_Start     = ""
	Room_End       = ""
	the_rooms      = make(map[string][]string)
	visit          = make(map[string]bool)
	RestOfAnt      = 0
	tunnels        = [][]string{}

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
		fmt.Println(tunnels)
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

func Split_Char(str string, char rune) int {
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
	if the_rooms[(Firstroom)] == nil {
		neighbr := []string{(neighbors)}
		the_rooms[(Firstroom)] = neighbr
	} else {
		the_rooms[(Firstroom)] = append(the_rooms[(Firstroom)], (neighbors))
	}
	if the_rooms[(neighbors)] == nil {
		neighbr := []string{(Firstroom)}
		the_rooms[(neighbors)] = neighbr
	} else {
		the_rooms[(neighbors)] = append(the_rooms[(neighbors)], (Firstroom))
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
		first := false
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

				i := Split_Char(text, '-')
				if i != -1 {
					Relation_Room(string(text[:i]), string(text[i+1:]))
				}
				fmt.Println(text)
			}
		}
		// for key, strr := range the_rooms {
		// 	fmt.Printf("%s-%s ", key, strr)
		// }

		i := Split_Char((start[1]), ' ')
		Room_Start = start[1][:i]
		i = Split_Char((end[1]), ' ')
		Room_End = end[1][:i]
		fmt.Println(Room_Start, ",", Room_End)
		Ant_Path("1", "0")
		fmt.Println()
		for _, paths := range tunnels {
			fmt.Print(paths)
			fmt.Println()
		}

		fmt.Println(len(tunnels))
		Find_Short_Path()
		for i, path := range tunnels {
			fmt.Println(i, path)
		}
		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading file:", err)
			return
		}
	}
}
