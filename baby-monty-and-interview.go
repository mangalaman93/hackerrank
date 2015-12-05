package main

import (
	"bufio"
	"fmt"
	"os"
)

func visit(nodeid, level, lastnodeid int, nbr []map[int]bool, counts []int) {
	counts[level] += 1
	level = (level + 1) % 2
	for n, _ := range nbr[nodeid] {
		if n != lastnodeid {
			visit(n, level, nodeid, nbr, counts)
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var T int
	fmt.Fscan(reader, &T)
	for t := 0; t < T; t++ {
		var N int
		fmt.Fscan(reader, &N)

		nbr := make([]map[int]bool, N)
		for i := 0; i < N; i++ {
			nbr[i] = make(map[int]bool)
		}

		var a, b int
		for i := 0; i < N-1; i++ {
			fmt.Fscan(reader, &a, &b)
			nbr[a-1][b-1] = true
			nbr[b-1][a-1] = true
		}

		if N == 1 {
			fmt.Fprintln(writer, 0)
			continue
		}

		counts := make([]int, 2)
		visit(0, 0, -1, nbr, counts)
		fmt.Fprintln(writer, counts[0]*(counts[0]-1)/2+counts[1]*(counts[1]-1)/2)
	}
}
