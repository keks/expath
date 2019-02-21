// plain implements the WalkState for plain old Dijkstra
package plain

import (
	"github.com/keks/expath"
)

func New(n expath.Node, dist, maxDist int, blocks expath.Graph) expath.WalkState {
	return walkState{
		n:       n,
		dist:    dist,
		maxDist: maxDist,
	}
}

type walkState struct {
	maxDist, dist int
	n             expath.Node
}

func (ws walkState) Node() expath.Node {
	return ws.n
}

func (ws walkState) MinDist() int {
	return ws.dist
}

func (ws walkState) Reachable() bool {
	return ws.dist <= ws.maxDist
}

func (ws walkState) Next(n expath.Node) expath.WalkState {
	return walkState{
		n:       n,
		maxDist: ws.maxDist,
		dist:    ws.dist + 1,
	}
}

func (left walkState) Merge(right_ expath.WalkState) expath.WalkState {
	right := right_.(walkState)

	if left.dist < right.dist {
		return left
	}

	return right
}
