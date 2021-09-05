package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

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

	input1 := ""
	input2 := ""
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("Please enter your root: ")
		scanner.Scan()
		text := scanner.Text()
		if len(text) != 0 {
			// fmt.Println("Your root is: ", text)
			input1 = text
			for {
				fmt.Printf("Please enter your destination: ")
				scanner.Scan()
				text := scanner.Text()
				if len(text) != 0 {
					// fmt.Println("Your destination is: ", text)
					input2 = text
					break
				} else {
					fmt.Println("You have entered nothing! Please try again")
				}
			}
			break
		} else {
			fmt.Println("You have entered nothing! Please try again")
		}

	}

	// bfs_list := bfs(nodes[input]) // assuming that all graphs have "A"
	// for i := range bfs_list {
	// 	fmt.Printf("Node %s's Depth: %d\n", bfs_list[i].value, bfs_list[i].depth)
	// }

	fmt.Println("Path: ", findShortestPath(input1, input2, nodes),
		"\nDistance: ", nodes[input2].depth-nodes[input1].depth)
}
