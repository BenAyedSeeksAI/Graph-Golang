package main

import (
	"fmt"
)

func AdjancencyList(edges [][]int) map[int][]int {
	graph := make(map[int][]int)

	for _, edge := range edges {
		if _, OK := graph[edge[0]]; !OK {
			graph[edge[0]] = make([]int, 0)
		}
		if _, OK := graph[edge[1]]; !OK {
			graph[edge[1]] = make([]int, 0)
		}
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}
	return graph
}
func HasConnection(grph map[int][]int, visited []bool, source int, destination int) bool {
	if source == destination {
		return true
	}
	if visited[source] {
		return false
	}
	visited[source] = true
	for _, nearNode := range grph[source] {
		if HasConnection(grph, visited, nearNode, destination) {
			return true
		}
	}
	return false
}
func DepthFirstSearch(grph map[int][]int, visited []bool, at int) {
	if visited[at] {
		return
	}
	visited[at] = true
	fmt.Println("We are at node : ", at)
	neighbors := grph[at]
	for _, node := range neighbors {
		DepthFirstSearch(grph, visited, node)
	}
}
func ReturnEdges(neighbors [][]int) map[int]struct{} {
	edges := map[int]struct{}{}
	for _, node := range neighbors {
		edges[node[0]] = struct{}{}
		edges[node[1]] = struct{}{}
	}
	return edges
}
func Display(output map[int][]int) {
	fmt.Println("map displayed : ")
	fmt.Println(output)
}
func TestingConnection(grph map[int][]int) {
	var maxy int
	for key := range grph {
		if key > maxy {
			maxy = key
		}
	}
	for source := range grph {
		for destination := range grph {
			if destination != source {
				vis := make([]bool, maxy+1)
				if HasConnection(grph, vis, source, destination) {
					fmt.Printf("[%d, %d] : Connected\n", source, destination)
				} else {
					fmt.Printf("[%d, %d] : Disconnected\n", source, destination)
				}
			}

		}
	}
}
func main() {
	neighbors := [][]int{{0, 1}, {0, 5}, {0, 4}, {2, 6}, {6, 3}, {3, 9}}
	edges := ReturnEdges(neighbors)
	fmt.Println("edges :", edges, len(edges))
	output := AdjancencyList(neighbors)
	Display(output)
	TestingConnection(output)
	DepthFirstSearch(output, make([]bool, 10000), 9)
}
