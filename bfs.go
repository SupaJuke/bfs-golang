package main

import (
	"fmt"
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
	fmt.Println("Hello Planet!")
	nodes := makeNode('a', 'b', 'c', 'd', 'e')
	nodes[0].neighbor = []*Node{&nodes[1], &nodes[4]}
	nodes[1].neighbor = []*Node{&nodes[0], &nodes[2], &nodes[3], &nodes[4]}
	nodes[2].neighbor = []*Node{&nodes[1], &nodes[3]}
	nodes[3].neighbor = []*Node{&nodes[1], &nodes[2], &nodes[4]}
	nodes[4].neighbor = []*Node{&nodes[0], &nodes[3]}

	bfs_list := bfs(&nodes[0])
	for i := range bfs_list {
		fmt.Printf("%c: %d\n", bfs_list[i].value, bfs_list[i].depth)
	}
}
