package funchandler

import (
	"sort"
)

func Init_Path_Groups(start, end string, the_rooms map[string][]string) [][][]string {
	InitPath := [][][]string{}
	for _, nehbior := range the_rooms[start] {
		Best_Path := [][]string{}
		BufSlc := BfsFirstEnd(start, nehbior, end, the_rooms, map[string]bool{})
		if len(BufSlc) > 0 {
			Best_Path = append(Best_Path, BufSlc)
			InitPath = append(InitPath, Best_Path)
		}

	}
	for _, nehbior := range the_rooms[start] {
		Best_Path := [][]string{}
		BufSlc := Bfs(start, nehbior, end, the_rooms, map[string]bool{})
		if len(BufSlc) > 0 {
			Best_Path = append(Best_Path, BufSlc)
			InitPath = append(InitPath, Best_Path)
		}
	}
	return InitPath
}

func BfsFirstEnd(start, startneibor, end string, the_rooms map[string][]string, visit map[string]bool) []string {
	quene := []string{startneibor}
	parent := make(map[string]string)
	visit[startneibor] = true
	visit[start] = true

	for len(quene) > 0 {
		current := quene[0]
		quene = quene[1:]

		for _, neighbor := range the_rooms[current] {
			if !visit[neighbor] || neighbor == end {
				visit[neighbor] = true
				parent[neighbor] = current
				quene = append(quene, neighbor)
			}
			if neighbor == end {
				return Complete_Path(parent, startneibor, end)
			}
		}
	}
	return nil
}

func Bfs(start, startneibor, end string, the_rooms map[string][]string, visit map[string]bool) []string {
	quene := []string{startneibor}
	parent := make(map[string]string)
	visit[startneibor] = true
	visit[start] = true

	for len(quene) > 0 {
		current := quene[0]
		quene = quene[1:]
		if current == end {
			return Complete_Path(parent, startneibor, end)
		}
		for _, neighbor := range the_rooms[current] {
			if !visit[neighbor] || neighbor == end {
				visit[neighbor] = true
				parent[neighbor] = current
				quene = append(quene, neighbor)
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
	// reverse
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
	return path
}

func Get_All_Path(start, end string, the_rooms map[string][]string, tunnels [][][]string) [][][]string {
	for i, Group_Path := range tunnels {
		visit := MarkVisist(Group_Path)
		for _, nehbior := range the_rooms[start] {
			if !visit[nehbior] {
				Paths := Bfs(start, nehbior, end, the_rooms, visit)
				if len(Paths) > 0 {
					tunnels[i] = append(tunnels[i], Paths)
					visit = MarkVisist(tunnels[i])
				}
			}
		}
	}
	for _, Group_Path := range tunnels {
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
