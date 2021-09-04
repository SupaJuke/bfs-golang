package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Only works for connected graph
func bfs(root *Node) []*Node {
	root.depth = 0
	gray := enqueue([]*Node{}, root)
	black := []*Node{}
	var node *Node
	for len(gray) > 0 {
		// fmt.Println("gray before dequeue:", gray)
		gray, node = dequeue(gray)
		// fmt.Println("gray after dequeue:", gray)
		for i := 0; i < len(node.neighbor); i++ {
			// the node was already discovered
			if node.neighbor[i].depth != -1 {
				continue
			}
			node.neighbor[i].parent = node
			node.neighbor[i].depth = node.depth + 1
			gray = enqueue(gray, node.neighbor[i])
			// fmt.Println("gray:", gray)
		}
		black = enqueue(black, node)
		// fmt.Println("black", black)
	}
	return black
}

func main() {

	/* Opening and reading from a file */
	filename := os.Args[1]
	fmt.Println("Filename: ", filename)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Path Error: ", err)
		os.Exit(1)
	}
	defer file.Close()

	var nodes map[string]*Node
	adjList := map[string][]string{}
	reader := bufio.NewScanner(file)
	for reader.Scan() {
		// Accessing out-neighborhood of each node
		// Printing each line that you read
		if node := strings.Split(reader.Text(), ":"); len(node) != 1 {
			adjList[node[0]] = strings.Split(node[1], " ")
		} else {
			// Reading the header: will generate all nodes with the function
			nodes = makeNode(strings.Split(reader.Text(), " "))
		}
	}

	for k := range adjList {
		fmt.Printf("Node%s -> %v\n", k, adjList[k])
	}

	if err := reader.Err(); err != nil {
		log.Fatal("File Error: ", err)
	}

	connectNodes(adjList, nodes)

	bfs_list := bfs(nodes["A"]) // assuming that all graphs have "A"
	for i := range bfs_list {
		fmt.Printf("Node %s's Depth: %d\n", bfs_list[i].value, bfs_list[i].depth)
	}
}
