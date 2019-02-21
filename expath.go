package expath // import "github.com/keks/expath"

import (
	"container/heap"
)

type Node interface {
	Label() string
}

type Nodes interface {
	Next() bool
	Get() Node
	Len() int
}

type Graph interface {
	Nodes() Nodes
	Outbound(Node) Nodes
}

type WalkState interface {
	Reachable() bool
	Next(n Node) WalkState
	Merge(WalkState) WalkState
	MinDist() int
	Node() Node
}

func MergeNext(wsSrc, wsDst WalkState, n Node) WalkState {
	return wsDst.Merge(wsSrc.Next(n))
}

type Path []Node

type NewWalkState func(n Node, dist, maxDist int, blocks Graph) WalkState

func Do(follows,blocks Graph, src, dst Node, max int, newWalkState NewWalkState) WalkState {
	var (
		// for all practical purposes, infinity is more than the maximum path length
		inf = max + 1
		ns  = follows.Nodes()
		Q   = newQueue(0, ns.Len())
	)

	for ns.Next() {
		var (
			n  = ns.Get()
			ws = newWalkState(n, inf, max, blocks)
		)

		heap.Push(Q, &item{ws: ws})
	}

	Q.update(src, newWalkState(src, 0, max, blocks))

	for Q.Len() > 0 {
		var (
			item = heap.Pop(Q).(*item)
			u    = item.ws.Node()
		)

		// we found the shortest path to our target, done!
		if u == dst {
			break
		}

		// Abort if we get a node with too long distance. We read
		// from a queue; next reads will not return closer nodes.
		if Q.WalkState(u).MinDist() > max {
			break
		}

		ns := follows.Outbound(u)
		for ns.Next() {
			var (
				n   = ns.Get()
				uWs = Q.WalkState(u)
				nWs = Q.WalkState(n)
			)

			Q.update(n, MergeNext(uWs, nWs, ns.Get()))
		}
	}

	return Q.WalkState(dst)
}
