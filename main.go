package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

var (
	start      = []string{}
	end        = []string{}
	Room_Start = ""
	Room_End   = ""
	the_rooms  = make(map[string][]string)
	RestOfAnt  = 0

	number_of_ants = -1
)

func Find_Best_Group(AllGroup [][][]string) ([][]string, int, int) {
	turn := 0
	IndexOfBestGroup := -1
	for i, group := range AllGroup {

		Ant := number_of_ants
		for Ant != 0 {
			Ant--
			group[0] = append(group[0], "L"+strconv.Itoa(number_of_ants-Ant))
			if len(group) > 1 && len(group[0]) > len(group[1]) {
				break
			}
		}

		for Ant != 0 {
			for i := range group {
				if Ant == 0 {
					break
				}
				if i < len(group)-1 && len(group[i]) > len(group[i+1]) {
					Ant--
					group[i+1] = append(group[i+1], "L"+strconv.Itoa(number_of_ants-Ant))
				} else if i < len(group)-1 && len(group[i]) == len(group[i+1]) {
					Ant--
					group[i] = append(group[i], "L"+strconv.Itoa(number_of_ants-Ant))
				}

			}
		}
		if turn == 0 || len(group[0]) < turn {
			turn = len(group[0])

			IndexOfBestGroup = i
		}
	}
	return AllGroup[IndexOfBestGroup], IndexOfBestGroup, turn - 1
}

func Move_Ant() {
	Tunnels, Indx_Grp, Turn := Find_Best_Group(Get_All_Path(Room_Start, Room_End, Init_Path_Groups(Room_Start, Room_End)))
	Base_Tunnels := Get_All_Path(Room_Start, Room_End, Init_Path_Groups(Room_Start, Room_End))
	Bs_Tunnel := Base_Tunnels[Indx_Grp]
	fmt.Println(Bs_Tunnel)
	// fmt.Println(Tunnels)

	for i := range Tunnels {
		Tunnels[i] = Tunnels[i][len(Bs_Tunnel[i]):]
	}
	fmt.Println(Tunnels)
	if Turn == 0 {
		Turn = 0
	}

	// Ant_key := make(map[string][]string)
	// for i, path := range Tunnels {
	// 	for _, room := range path {
	// 		for j, path2 := range Bs_Tunnel {
	// 			if i == j {
	// 				Ant_key[room] = path2
	// 			}
	// 		}
	// 	}
	// }
	Turn2 := Turn

	Ant_Position := make(map[string]int)
	for _, path := range Tunnels {
		for _, room := range path {
			Ant_Position[room] = 0
		}
	}
	fmt.Println(Ant_Position)
	for Turn != 0 {
		i := 0
		for i < len(Tunnels) {
			j := 0
			for j < len(Tunnels[i]) && j <= Turn2-Turn {
				if Ant_Position[Tunnels[i][j]] < len(Bs_Tunnel[i]) {
					fmt.Print(Tunnels[i][j] + "-" + Bs_Tunnel[i][Ant_Position[Tunnels[i][j]]] + " ")
					Ant_Position[Tunnels[i][j]] = Ant_Position[Tunnels[i][j]] + 1
				}
				j++

			}
			i++

		}
		Turn--
		fmt.Println()

	}
	// fmt.Println(Ant_key)
	// fmt.Print(Tunnels[Indx_Grp][0][len(Base_Tunnels[Indx_Grp][0])],"-",Base_Tunnels[Indx_Grp][0][0])
}

func Init_Path_Groups(start, end string) [][][]string {
	tunnels := [][][]string{}
	for _, nehbior := range the_rooms[start] {
		Best_Path := [][]string{}
		Best_Path = append(Best_Path, Bfs(nehbior, end, map[string]bool{start: true, nehbior: true}))
		tunnels = append(tunnels, Best_Path)
	}
	return tunnels
}

func Bfs(start, end string, visit map[string]bool) []string {
	quene := []string{start}
	parent := make(map[string]string)
	visit[start] = true
	visit[Room_Start] = true
	for len(quene) > 0 {
		current := quene[0]
		quene = quene[1:]

		for _, neighbor := range the_rooms[current] {
			if !visit[neighbor] || neighbor == end {
				visit[neighbor] = true
				parent[neighbor] = current
				quene = append(quene, neighbor)
			}
			if current == end {
				return Complete_Path(parent, start, end)
			}
		}
	}
	return nil
}

func Complete_Path(parent map[string]string, start, end string) []string {
	path := []string{}
	curnt := end
	for curnt != start {

		path = append(path, curnt)

		curnt = parent[curnt]
	}
	path = append(path, start)
	// path = append(path, Room_Start)
	slices.Reverse(path)
	return path
}

func Get_All_Path(start, end string, tunnels [][][]string) [][][]string {
	for i, Group_Path := range tunnels {

		visit := MarkVisist(Group_Path)
		for _, nehbior := range the_rooms[start] {
			if !visit[nehbior] {
				Paths := Bfs(nehbior, end, visit)
				if len(Paths) > 0 {
					tunnels[i] = append(tunnels[i], Paths)
					visit = MarkVisist(tunnels[i])
				}

			}
		}
		sort.Slice(Group_Path, func(i, j int) bool {
			return len(Group_Path[i]) < len(Group_Path[j])
		})
	}

	return tunnels
}

func MarkVisist(Group_Path [][]string) map[string]bool {
	visit := map[string]bool{}
	for _, paths := range Group_Path {
		for _, Room := range paths {
			visit[Room] = true
		}
	}
	return visit
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
		parseFile(args[0])
		i := 0
		fmt.Println(end)
		for _, ValueEnd := range end {
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

		fmt.Println("          OK")

		fmt.Println("          OK")
		// fmt.Println(Get_All_Path(Room_Start, Room_End, Init_Path_Groups(Room_Start, Room_End)))

		// // fmt.Println(start)
		// // fmt.Println(end)
		// fmt.Println(Find_Best_Group(Get_All_Path(Room_Start, Room_End, Init_Path_Groups(Room_Start, Room_End))))
		// fmt.Println(Get_All_Path(Room_Start, Room_End, Init_Path_Groups(Room_Start, Room_End)))
		Move_Ant()

	}
}

func parseFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

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
}
