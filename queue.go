package main

/* Node Part */
type Node struct {
	value    string
	parent   *Node
	neighbor []*Node
	depth    int
}

func makeNode(chars []string) map[string]*Node {
	nodes := map[string]*Node{}
	for _, char := range chars {
		nodes[char] = &Node{
			value:    char,
			parent:   nil,
			neighbor: []*Node{},
			depth:    -1,
		}
	}
	return nodes
}

func connectNodes(adjList map[string][]string, nodes map[string]*Node) {
	for node, outNodes := range adjList { // map string to slice of string
		for _, outNode := range outNodes { // slice of string
			nodes[node].neighbor = append(nodes[node].neighbor, nodes[outNode])
		}
	}
}

////////////////////////////////////////////////////////////////////////

/* Queue Part: Slice of Pointers of Nodes */

// Insert a pointer to a node to the bottom of the queue
func enqueue(queue []*Node, node *Node) []*Node {
	return append(queue, node)
	// queue[len(queue)] = node
}

// Remove & Return a pointer to a node in front of the queue
func dequeue(queue []*Node) ([]*Node, *Node) {
	var temp = queue[0]
	queue = queue[1:]
	return queue, temp
}
