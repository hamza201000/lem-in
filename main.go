package main

import (
	"fmt"
	"os"

	"lemin/funchandler"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Usage: go run . input.txt")
		return
	}

	graph, err := funchandler.ParseFileToGraph(args[0])
	if err != nil {
		fmt.Println(err)
		return
	}

	// set global ants and start/end names

	if graph.Start == nil || graph.End == nil {
		fmt.Println("ERROR: start or end not set by parser")
		return
	}

	// now run algorithm
	fmt.Println(graph.Ants)
	fmt.Println(graph.Start)
	Base_Tunnels := funchandler.Get_All_Path(graph.Start.Name, graph.End.Name, graph.The_rooms, (funchandler.Init_Path_Groups(graph.Start.Name, graph.End.Name, graph.The_rooms)))
	fmt.Println(Base_Tunnels)
	Tunnels, Indx_Grp, Turn := funchandler.Find_Best_Group(graph.Ants, funchandler.CopySlice(Base_Tunnels))

	Bs_Tunnel := Base_Tunnels[Indx_Grp]

	funchandler.Move_Ant(graph.Ants, graph.End.Name, Tunnels, Bs_Tunnel, Turn)
}
