package main

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

func findShortestPath(u string, v string, graph map[string]*Node) string {
	root := graph[u]
	// running bfs on u
	bfs(root)
	node := graph[v]
	path := v
	for node.parent != nil {
		node = node.parent
		path = node.value + " -> " + path
	}
	return path
}
