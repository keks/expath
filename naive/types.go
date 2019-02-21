package naive

import (
	"github.com/keks/expath"
)

type StringNode string

func (n StringNode) Label() string {
	return string(n)
}

type SliceNodes []expath.Node

// insert leading nil
func NewSliceNodes(ins ...expath.Node) expath.Nodes {
	ns := make(SliceNodes, len(ins)+1)
	copy(ns[1:], ins)
	return &ns
}

func (ns *SliceNodes) Len() int {
	return len(*ns) - 1
}

func (ns *SliceNodes) Next() bool {
	if len(*ns) == 0 {
		return false
	}

	*ns = (*ns)[1:]
	return len(*ns) > 0
}

func (ns *SliceNodes) Get() expath.Node {
	return (*ns)[0]
}

type Graph struct {
	V expath.Nodes
	E map[expath.Node]expath.Nodes
}

func (g *Graph) Nodes() expath.Nodes {
	return g.V
}

func (g *Graph) Outbound(n expath.Node) expath.Nodes {
	ns, ok := g.E[n]
	if !ok {
		ns = &SliceNodes{nil}
	}

	return ns
}

func MakeGraph(pairs ...StringNode) expath.Graph {
	V := extractNodes(pairs...)
	E := extractEdges(pairs...)

	return &Graph{V: V, E: E}
}

func extractNodes(pairs ...StringNode) expath.Nodes {
	// first element will be consumed by next call
	var ns = SliceNodes{nil}
	has := make(map[string]struct{})

	for _, n := range pairs {
		if _, ok := has[n.Label()]; ok {
			continue
		}

		has[n.Label()] = struct{}{}
		ns = append(ns, n)
	}

	return &ns
}

func extractEdges(pairs ...StringNode) map[expath.Node]expath.Nodes {
	if len(pairs)%2 != 0 {
		return nil
	}

	type edge struct{ src, dst string }

	pairToEdge := func(src, dst expath.Node) edge {
		return edge{
			src: src.Label(),
			dst: dst.Label(),
		}
	}

	getPair := func(pairs []StringNode) (expath.Node, expath.Node, []StringNode) {
		from := pairs[0]
		to := pairs[1]
		pairs = pairs[2:]

		return from, to, pairs
	}

	has := make(map[edge]struct{})
	E := make(map[expath.Node]expath.Nodes)

	var src, dst expath.Node
	for len(pairs) > 0 {
		src, dst, pairs = getPair(pairs)

		e := pairToEdge(src, dst)
		if _, ok := has[e]; ok {
			continue
		}
		has[e] = struct{}{}

		ns := E[src]
		if ns == nil {
			ns = &SliceNodes{nil}
		}

		ns_ := append(*(ns.(*SliceNodes)), dst)
		E[src] = &ns_
	}

	return E
}
