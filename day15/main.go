package main

import (
	"container/heap"
	"fmt"
	"math"
	"strconv"

	"github.com/dotMaro/AoC2021/utils"
)

func main() {
	input := utils.InputString("day15/input.txt")
	riskMap := parseRiskMap(input)
	fmt.Println(riskMap.extend().dijkstrasPriority())
}

const (
	width  = 30
	height = 30
)

type riskMap [][]uint8

type coord struct {
	x, y int
}

func parseRiskMap(s string) riskMap {
	lines := utils.SplitLine(s)
	riskMap := make([][]uint8, len(lines))
	for y, line := range lines {
		riskMap[y] = make([]uint8, len(line))
		for x, r := range line {
			value, _ := strconv.Atoi(string(r))
			riskMap[y][x] = uint8(value)
		}
	}
	return riskMap
}

func (r riskMap) extend() riskMap {
	extended := make([][]uint8, len(r)*5)
	for y := range extended {
		extended[y] = make([]uint8, len(r[0])*5)
	}
	for tileY := 0; tileY < 5; tileY++ {
		for tileX := 0; tileX < 5; tileX++ {
			for y, row := range r {
				for x, risk := range row {
					extended[tileY*len(r)+y][tileX*len(r[0])+x] = risk + uint8(tileX) + uint8(tileY)
				}
			}
		}
	}
	return extended
}

func (r riskMap) neighbors(c coord) []coord {
	var neighbors []coord
	if c.x > 0 {
		neighbors = append(neighbors, coord{x: c.x - 1, y: c.y})
	}
	if c.x < len(r[0])-1 {
		neighbors = append(neighbors, coord{x: c.x + 1, y: c.y})
	}
	if c.y > 0 {
		neighbors = append(neighbors, coord{x: c.x, y: c.y - 1})
	}
	if c.y < len(r)-1 {
		neighbors = append(neighbors, coord{x: c.x, y: c.y + 1})
	}
	return neighbors
}

func (r riskMap) h(c coord) int {
	return len(r) - c.y + len(r[0]) - c.x
}

func (r riskMap) reconstructPathCost(cameFrom map[coord]coord, current coord) int {
	// totalPath := 0
	totalCost := 0

	for current := range cameFrom {
		totalCost += int(r[current.x][current.y])
	}
	return totalCost
}

func (r riskMap) aStar(start, goal coord) int {
	openSet := map[coord]struct{}{
		start: {},
	}

	cameFrom := make(map[coord]coord)

	gScore := map[coord]int{
		start: 0,
	}

	fScore := map[coord]int{
		start: r.h(start), // ??
	}

	for len(openSet) != 0 {
		var lowestScoreNode coord
		lowestScore := math.MaxUint32
		for node := range openSet {
			score, hasScore := fScore[node]
			if hasScore && score < lowestScore {
				lowestScoreNode = node
				lowestScore = score
			}
		}
		current := lowestScoreNode
		if current == goal {
			return r.reconstructPathCost(cameFrom, current)
		}

		delete(openSet, current)
		for _, neighbor := range r.neighbors(current) {
			tentativeGScore := gScore[current] + int(r[neighbor.y][neighbor.x])
			if tentativeGScore < gScore[neighbor] {
				cameFrom[neighbor] = current
				gScore[neighbor] = tentativeGScore
				fScore[neighbor] = tentativeGScore + r.h(neighbor)
				_, neighborInOpenSet := openSet[neighbor]
				if !neighborInOpenSet {
					openSet[neighbor] = struct{}{}
				}
			}
		}
	}
	return 0
}

func (r riskMap) dijkstras() uint {
	q := make(map[coord]struct{})
	dist := make(map[coord]uint)
	prev := make(map[coord]coord)

	for y, row := range r {
		for x := range row {
			coord := coord{x: x, y: y}
			dist[coord] = math.MaxUint32
			// prev[coord]
			q[coord] = struct{}{}
		}
	}
	dist[coord{y: 0, x: 0}] = 0

	for len(q) > 0 {
		var u coord
		var minDist uint = math.MaxUint32
		for c := range q {
			curDist := dist[c]
			if curDist < minDist {
				minDist = curDist
				u = c
			}
		}

		delete(q, u)

		if u.x == len(r[0])-1 && u.y == len(r)-1 {
			var totalCost uint = 0
			for !(u.x == 0 && u.y == 0) {
				// fmt.Printf("%d (%d,%d) prev %v\n", uint(r[u.y][u.x]), u.x, u.y, prev[u])
				totalCost += uint(r[u.y][u.x])
				u = prev[u]
			}
			return totalCost
		}

		for _, neighbor := range r.neighbors(u) {
			_, inQ := q[neighbor]
			if !inQ {
				continue
			}
			alt := dist[u] + uint(r[neighbor.y][neighbor.x])
			if alt < dist[neighbor] {
				dist[neighbor] = alt
				prev[neighbor] = u
			}
		}
	}
	return 0
}

// 7617 too high

func (r riskMap) dijkstrasPriority() uint {
	dist := make(map[coord]uint)
	prev := make(map[coord]coord)

	q := make(PriorityQueue, len(r)*len(r[0]))

	dist[coord{y: 0, x: 0}] = 0
	i := 0
	for y, row := range r {
		for x := range row {
			coord := coord{x: x, y: y}
			if x != 0 && y != 0 {
				dist[coord] = math.MaxUint32
			}
			q[i] = &Item{
				value:    coord,
				priority: int(dist[coord]),
				index:    i,
			}
			i++
		}
	}
	heap.Init(&q)

	for len(q) > 0 {
		u := heap.Pop(&q).(*Item)
		uNode := u.value

		prevNode := prev[uNode]
		if prevNode.x != 0 && prevNode.y != 0 && uNode.x == len(r[0])-1 && uNode.y == len(r)-1 {
			var totalCost uint = 0
			for !(uNode.x == 0 && uNode.y == 0) {
				fmt.Printf("%d (%d,%d) prev %v\n", uint(r[uNode.y][uNode.x]), uNode.x, uNode.y, prev[uNode])
				totalCost += uint(r[uNode.y][uNode.x])
				uNode = prev[uNode]
			}
			return totalCost
		}

		for _, neighbor := range r.neighbors(u.value) {
			alt := dist[u.value] + uint(r[neighbor.y][neighbor.x])
			if alt < dist[neighbor] {
				dist[neighbor] = alt
				prev[neighbor] = u.value
				q.update(u, u.value, u.priority-1)
			}
		}
	}
	return 0
}

// Below was shamelessly copied from an example in the container/heap docs.

// An Item is something we manage in a priority queue.
type Item struct {
	value    coord // The value of the item; arbitrary.
	priority int   // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// Lowest to highest.
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

// Push an item.
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

// Pop an item.
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value coord, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}
