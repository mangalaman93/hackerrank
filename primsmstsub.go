package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Edge struct {
	from, to, weight int
}

type Edges []*Edge

func (e Edges) Len() int {
	return len(e)
}

func (e Edges) Less(i, j int) bool {
	return e[i].weight < e[j].weight
}

func (e Edges) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func NewEdge(from, to, weight int) *Edge {
	return &Edge{from: from, to: to, weight: weight}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	var T int = 1

	for t := 0; t < T; t++ {
		var V, E int
		fmt.Fscan(reader, &V, &E)

		edges := make(Edges, 0)
		for i := 0; i < E; i++ {
			var x, y, r int
			fmt.Fscan(reader, &x, &y, &r)
			x = x - 1
			y = y - 1
			edges = append(edges, NewEdge(x, y, r))
		}
		sort.Sort(edges)

		var S int
		fmt.Fscan(reader, &S)
		S = S - 1

		num_nodes := 0
		total_weight := 0
		intree := make([]bool, V)
		for _, e := range edges {
			if !intree[e.from] || !intree[e.to] {
				total_weight += e.weight
				if !intree[e.from] {
					intree[e.from] = true
					num_nodes++
				}
				if !intree[e.to] {
					intree[e.to] = true
					num_nodes++
				}
			}

			if num_nodes == V {
				break
			}
		}

		fmt.Fprintln(writer, total_weight)
	}

	writer.Flush()
}
