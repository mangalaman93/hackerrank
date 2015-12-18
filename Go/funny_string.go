package main

import "fmt"

func abs(n int) int {
	if n < 0 {
		n = -n
	}

	return n
}

func main() {
	var N int
	fmt.Scan(&N)
	for j := 0; j < N; j++ {
		var s string
		fmt.Scan(&s)

		flag := true
		l := len(s)
		for i := 1; i < l; i++ {
			x := abs(int(s[i]) - int(s[i-1]))
			y := abs(int(s[l-1-i]) - int(s[l-i]))
			if x != y {
				flag = false
				break
			}
		}

		if flag {
			fmt.Println("Funny")
		} else {
			fmt.Println("Not Funny")
		}
	}
}
