package blocks

import (
	"github.com/keks/expath"
)

func New(n expath.Node, dist, maxDist int, blocks expath.Graph) expath.WalkState {
	ws:= walkState{
		n: n,
		maxDist: maxDist,
		blockSets: blockSets{dist: make(nodeSet)},
		blocks: blocks,
	}

	ws.blockSets.Truncate(maxDist)

	return ws
}

type nodeSet map[expath.Node]struct{}

func (ns nodeSet) Contains(n expath.Node) bool {
	_, ok := ns[n]
	return ok
}

func (ns nodeSet) Set(n expath.Node) {
	ns[n] = struct{}{}
}

func (left nodeSet) Intersect(right nodeSet) nodeSet {
	if right == nil {
		return left
	}

	if left == nil {
		return right
	}

	out := make(nodeSet)

	for n := range left {
		if right.Contains(n) {
			out.Set(n)
		}
	}

	return out
}

type blockSets map[int]nodeSet

func (left blockSets) Intersect(right blockSets) blockSets {
	out := make(blockSets)

	for dist, bs := range left {
		out[dist] = bs.Intersect(right[dist])
	}

	for dist, bs := range right {
		if _, ok := out[dist]; ok {
			continue
		}

		out[dist] = bs
	}

	return out
}

func (bss blockSets) Truncate(max int)	 {
	for dist := range bss {
		if dist > max {
			delete(bss, dist)
		}
	}
}

type walkState struct {
	maxDist int
	n       expath.Node

	blockSets blockSets
	blocks expath.Graph
}

func (ws walkState) Node() expath.Node {
	return ws.n
}

func (ws walkState) MinDist() int {
	var min = ws.maxDist + 1

	for dist := range ws.blockSets {
		if dist < min {
			min = dist
		}
	}

	return min
}

func (ws walkState) Reachable() bool {
	for dist:=0; dist<=ws.maxDist; dist++ {
		bs := ws.blockSets[dist]

		if bs == nil {
			continue
		}

		if _, ok := bs[ws.n]; ok {
			continue
		}

		return true
	}

	return false
}

func (ws walkState) Next(n expath.Node) expath.WalkState {
	bss := make(blockSets)
	for dist, bs := range ws.blockSets {
		if _, ok := bs[n]; ok {
			continue
		}

		bss[dist+1] = bs

		blocked := ws.blocks.Outbound(n)
		for blocked.Next() {
			bss[dist+1].Set(blocked.Get())
		}
	}

	nextWs := walkState{
		maxDist:   ws.maxDist,
		n:         n,
		blockSets: bss,
		blocks:    ws.blocks,
	}

	nextWs.blockSets.Truncate(ws.maxDist)

	return nextWs
}

func (left walkState) Merge(right_ expath.WalkState) expath.WalkState {
	right := right_.(walkState)

	nextWs := walkState{
		maxDist: left.maxDist,
		n: left.n,
		blockSets: left.blockSets.Intersect(right.blockSets),
		blocks: left.blocks,
	}

	return nextWs
}
