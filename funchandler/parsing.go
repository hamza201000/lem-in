package funchandler

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Room struct {
	Name      string
	X, Y      int
	Neighbors []*Room
	IsStart   bool
	IsEnd     bool
}

type Graph struct {
	The_rooms map[string][]string
	Rooms     map[string]*Room
	Start     *Room
	End       *Room
	Ants      int
}

// ParseFileToGraph parses the given filename into a Graph.
// It returns a detailed error when the input violates any expected rule.
func ParseFileToGraph(filename string) (*Graph, error) {
	f, err := os.Open(filename)                                          
	if err != nil {
		return nil, fmt.Errorf("ERROR: cannot open file: %v", err)
	}
	defer f.Close()

	graph := &Graph{Rooms: make(map[string]*Room), The_rooms: make(map[string][]string)}

	scanner := bufio.NewScanner(f)
	lineNumber := 0

	var antsRead bool
	var startNext bool
	var endNext bool
	var startFound bool
	var endFound bool
	var startLine, endLine int

	// used to detect duplicate coordinates
	coords := make(map[[2]int]string)

	for scanner.Scan() {
		lineNumber++
		raw := scanner.Text()
		line := strings.TrimSpace(raw)

		if line == "" {
			continue
		}

		// commands (must be after ants)
		if strings.HasPrefix(line, "##") {
			if !antsRead {
				return nil, fmt.Errorf("ERROR: invalid data format, command %s before number of ants (line %d)", line, lineNumber)
			}
			if line == "##start" {
				if startFound || startNext {
					return nil, fmt.Errorf("ERROR: invalid data format, duplicate ##start (line %d)", lineNumber)
				}
				startNext = true
				continue
			} else if line == "##end" {
				if endFound || endNext {
					return nil, fmt.Errorf("ERROR: invalid data format, duplicate ##end (line %d)", lineNumber)
				}
				endNext = true
				continue
			} else {
				return nil, fmt.Errorf("ERROR: invalid data format, unknown command %s (line %d)", line, lineNumber)
			}
		}

		if strings.HasPrefix(line, "#") {
			continue
		}

		// read ants (first non-comment non-empty line)
		if !antsRead {
			n, err := strconv.Atoi(line)
			if err != nil || n <= 0 {
				return nil, fmt.Errorf("ERROR: invalid data format, invalid number of ants (must appear first) (line %d)", lineNumber)
			}
			graph.Ants = n
			antsRead = true
			continue
		}

		parts := strings.Fields(line)
		if len(parts) == 3 {
			name := parts[0] /*  */

			if strings.HasPrefix(name, "L") || strings.HasPrefix(name, "#") {
				return nil, fmt.Errorf("ERROR: invalid data format, invalid room name '%s' (cannot start with 'L' or '#') (line %d)", name, lineNumber)
			}

			if _, exists := graph.Rooms[name]; exists {
				return nil, fmt.Errorf("ERROR: invalid data format, duplicate room name: %s (line %d)", name, lineNumber)
			}

			x, errX := strconv.Atoi(parts[1])
			y, errY := strconv.Atoi(parts[2])
			if errX != nil || errY != nil {
				return nil, fmt.Errorf("ERROR: invalid data format, invalid room coordinates: %s (line %d)", line, lineNumber)
			}

			// duplicate coordinates check (feature from second code)
			coordKey := [2]int{x, y}
			if existing, ok := coords[coordKey]; ok {
				return nil, fmt.Errorf("ERROR: Duplicate coordinates for rooms '%s' and '%s' (line %d)", existing, name, lineNumber)
			}
			coords[coordKey] = name

			room := &Room{Name: name, X: x, Y: y}

			if startNext {
				room.IsStart = true
				startLine = lineNumber
				graph.Start = room
				startNext = false
				startFound = true
			}
			if endNext {
				room.IsEnd = true
				endLine = lineNumber
				graph.End = room
				endNext = false
				endFound = true
			}

			if startLine > 0 && endLine > 0 {
				if startLine >= endLine {
					return nil, fmt.Errorf("ERROR: invalid data format, start room must be defined before end room (start at line %d, end at line %d)", startLine, endLine)
				}
			}

			graph.Rooms[name] = room

			continue
		}

		if strings.Contains(line, "-") {
			linkParts := strings.Split(line, "-")
			if len(linkParts) != 2 || linkParts[0] == "" || linkParts[1] == "" {
				return nil, fmt.Errorf("ERROR: invalid data format, invalid link line: %s (line %d)", line, lineNumber)
			}
			a := linkParts[0]
			b := linkParts[1]
			roomA, okA := graph.Rooms[a]
			roomB, okB := graph.Rooms[b]
			if !okA || !okB {
				return nil, fmt.Errorf("ERROR: invalid data format, link references unknown room: %s (line %d)", line, lineNumber)
			}

			for _, link := range roomA.Neighbors {
				if link == roomB {
					return nil, fmt.Errorf("ERROR: duplicate tunnel between %s and %s (line %d)", a, b, lineNumber)
				}
			}

			// append normally
			if graph.The_rooms[string(roomA.Name)] == nil {
				neighbr := []string{string(roomB.Name)}
				graph.The_rooms[string(roomA.Name)] = neighbr
			} else {
				graph.The_rooms[string(roomA.Name)] = append(graph.The_rooms[string(roomA.Name)], string(roomB.Name))
			}
			if graph.The_rooms[string(roomB.Name)] == nil {
				neighbr := []string{string(roomA.Name)}
				graph.The_rooms[string(roomB.Name)] = neighbr
			} else {
				graph.The_rooms[string(roomB.Name)] = append(graph.The_rooms[string(roomB.Name)], string(roomA.Name))
			}

			
			continue
		}

		return nil, fmt.Errorf("ERROR: invalid data format, invalid line: %s (line %d)", line, lineNumber)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("ERROR: reading file: %v", err)
	}

	if !antsRead {
		return nil, fmt.Errorf("ERROR: invalid data format, invalid number of ants")
	}
	if !startFound {
		return nil, fmt.Errorf("ERROR: invalid data format, no start room found")
	}
	if !endFound {
		return nil, fmt.Errorf("ERROR: invalid data format, no end room found")
	}

	return graph, nil
}
