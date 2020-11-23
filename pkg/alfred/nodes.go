package alfred

import "sync"

type NodePool struct {
	nodes map[string]NodeInterface
	w     sync.RWMutex
}

func NewNodePool() *NodePool {
	return &NodePool{
		nodes: make(map[string]NodeInterface),
	}
}

func (n *NodePool) AddNode(i NodeInterface) {
	n.w.Lock()
	defer n.w.Unlock()
	n.nodes[i.Sha()] = i
}

func (n *NodePool) RemoveNode(i NodeInterface) {
	n.w.Lock()
	defer n.w.Unlock()
	delete(n.nodes, i.Sha())
}

func (n *NodePool) UpdateNode(i NodeInterface) {
	n.AddNode(i)
}
