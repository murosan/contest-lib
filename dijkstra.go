package cutil

import (
	"container/heap"
	"math"
)

//start:Dijkstra:
//imports[math,container/heap]
type EdgeD struct {
	From     int
	To       int
	Distance int
}

type Cost struct {
	to       int
	priority int
	index    int
}

func Dijkstra(edges []EdgeD, start, goal int) (connected bool, cost int) {
	if len(edges) == 0 {
		return
	}

	// initialize
	max := math.MaxInt32
	distances := make(map[int]int)
	for _, e := range edges {
		distances[e.From] = max
		distances[e.To] = max
	}

	pq := make(EdgePriorityQueue, 0)
	heap.Init(&pq)

	// solve
	distances[start] = 0
	heap.Push(&pq, &Cost{to: start, priority: 0})

	for pq.Len() > 0 {
		cost := heap.Pop(&pq).(*Cost)
		if distances[cost.to] < cost.priority {
			continue
		}
		for _, e := range edges {
			c := distances[e.From] + e.Distance
			if distances[e.From] != max && distances[e.To] > c {
				distances[e.To] = c
				heap.Push(&pq, &Cost{to: e.To, priority: c})
			}
		}
	}

	if distances[goal] != max {
		connected = true
	}
	cost = distances[goal]

	return
}

type EdgePriorityQueue []*Cost

func (pq EdgePriorityQueue) Len() int { return len(pq) }
func (pq EdgePriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}
func (pq EdgePriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}
func (pq *EdgePriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Cost)
	item.index = n
	*pq = append(*pq, item)
}
func (pq *EdgePriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}
func (pq *EdgePriorityQueue) update(item *Cost, to, distance int) {
	item.to = to
	item.priority = distance
	heap.Fix(pq, item.index)
}

//end:Dijkstra:
