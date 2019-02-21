package expath

// This file is heavily inspired by the PriorityQueue example in container/heap.

import (
	"container/heap"
)

type item struct {
	ws    WalkState
	index int
}

func newQueue(len, cap int) *queue {
	return &queue{
		items:   make([]*item, len, cap),
		nodeMap: make(map[Node]*item),
	}
}

type queue struct {
	items   []*item
	nodeMap map[Node]*item
}

func (q queue) Len() int {
	return len(q.items)
}

func (q queue) Less(i, j int) bool {
	return q.items[i].ws.MinDist() < q.items[j].ws.MinDist()
}

func (q queue) Swap(i, j int) {
	q.items[i], q.items[j] = q.items[j], q.items[i]

	q.items[i].index = i
	q.items[j].index = j
}

func (q *queue) Push(x interface{}) {
	n := len(q.items)

	item := x.(*item)
	item.index = n

	q.items = append(q.items, item)
	q.nodeMap[item.ws.Node()] = item
}

func (q *queue) Pop() interface{} {
	old := q.items
	n := len(old)

	it := old[n-1]
	it.index = -1

	q.items = old[:n-1]

	// don't remove it from the map, so we can still query the walkstate
	return it
}

func (q *queue) update(n Node, ws WalkState) {
	it := q.nodeMap[n]
	it.ws = ws
	heap.Fix(q, it.index)
}

func (q *queue) WalkState(n Node) WalkState {
	return q.nodeMap[n].ws
}
