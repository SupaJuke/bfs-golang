package main

type Node struct {
	value    byte
	parent   *Node
	neighbor []*Node
	depth    int
}

func makeNode(chars ...byte) []Node {
	nodes := []Node{}
	for _, c := range chars {
		nodes = append(nodes, Node{
			value:    c,
			parent:   nil,
			neighbor: nil,
			depth:    -1,
		})
	}
	return nodes
}

// Queue = slice of pointers of Nodes

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
