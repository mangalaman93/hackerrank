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
		var N int
		fmt.Fscan(reader, &N)

		flag := true
		for x := (N / 3); x >= 0; x-- {
			y := (N - 3*x) / 5
			if (3*x + 5*y) == N {
				flag = false
				for i := 0; i < 3*x; i++ {
					fmt.Fprint(writer, 5)
				}
				for i := 0; i < 5*y; i++ {
					fmt.Fprint(writer, 3)
				}

				fmt.Fprint(writer, "\n")
				break
			}
		}

		if flag {
			fmt.Fprintln(writer, -1)
		}
	}

	writer.Flush()
}
