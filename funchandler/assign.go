package funchandler

import (
	"fmt"
	"strconv"
)

func Find_Best_Group(number_of_ants int,AllGroup [][][]string) ([][]string, int, int) {
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
			for j := range group {
				if Ant == 0 {
					break
				}
				if j < len(group)-1 && len(group[j]) <= len(group[j+1]) {
					Ant--
					group[j] = append(group[j], "L"+strconv.Itoa(number_of_ants-Ant))
				}
				if len(group) > 1 && j == len(group)-1 && len(group[j]) < len(group[j-1]) {
					Ant--
					group[j] = append(group[j], "L"+strconv.Itoa(number_of_ants-Ant))
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

func Move_Ant(number_of_ants int,Room_End string,Tunnels, Bs_Tunnel [][]string, Turn int) {
	
	for i := range Tunnels {
		Tunnels[i] = Tunnels[i][len(Bs_Tunnel[i]):]
	}

	Turn2 := Turn

	Ant_Position := make(map[string]int)
	for _, path := range Tunnels {
		for _, room := range path {
			Ant_Position[room] = 0
		}
	}

	i := 0
	Terminal := false
	
	for !Terminal {
		i = 0
		for i < len(Tunnels) {
			j := 0
			for j < len(Tunnels[i]) && j <= Turn2-Turn {
				if Ant_Position[Tunnels[i][j]] < len(Bs_Tunnel[i]) {
					fmt.Print(Tunnels[i][j] + "-" + Bs_Tunnel[i][Ant_Position[Tunnels[i][j]]] + " ")
					nmb, err := strconv.Atoi(Tunnels[i][j][1:])
					if err != nil {
					}
					if nmb == number_of_ants && Bs_Tunnel[i][Ant_Position[Tunnels[i][j]]] == Room_End {
						Terminal = true
					}
					Ant_Position[Tunnels[i][j]] = Ant_Position[Tunnels[i][j]] + 1
				}
				j++
			}
			i++
		}
		
		Turn--
		fmt.Println()
	
	}
}
