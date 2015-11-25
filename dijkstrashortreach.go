package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	var T int
	fmt.Fscan(reader, &T)

	for t := 0; t < T; t++ {
		var V, E int
		fmt.Fscan(reader, &V, &E)

		adj := make([][]int, V)
		for i := 0; i < V; i++ {
			adj[i] = make([]int, V)
		}

		for i := 0; i < E; i++ {
			var x, y, r int
			fmt.Fscan(reader, &x, &y, &r)
			x = x - 1
			y = y - 1
			if adj[x][y] == 0 || adj[x][y] > r {
				adj[x][y] = r
				adj[y][x] = r
			}
		}

		var S int
		fmt.Fscan(reader, &S)
		S = S - 1
		distance := make([]int, V)
		for i := 0; i < V; i++ {
			distance[i] = -1
		}

		// algorithm
		distance[S] = 0
		visited := make([]int, V)
		for {
			mind := int((^uint(0)) >> 1)
			index := -1
			for i := 0; i < V; i++ {
				if visited[i] != 1 {
					if distance[i] != -1 && distance[i] < mind {
						mind = distance[i]
						index = i
					}
				}
			}

			if index == -1 {
				break
			}

			visited[index] = 1
			for i := 0; i < V; i++ {
				if adj[index][i] > 0 {
					if visited[i] != 1 {
						if distance[i] == -1 || distance[i] > (distance[index]+adj[index][i]) {
							distance[i] = distance[index] + adj[index][i]
						}
					}
				}
			}
		}

		for i := 0; i < V; i++ {
			if i != S {
				fmt.Fprint(writer, distance[i], " ")
			}
		}
		fmt.Fprintln(writer, "")
	}

	writer.Flush()
}
