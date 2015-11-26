package main

import (
	"bufio"
	"fmt"
	"os"
)


func binarySearch(v int, ar []int, begin, end int) int {
    mid := (begin+end)/2
    if mid == begin {
        return end
    }
    if v > ar[mid] && v <= ar[mid-1] {
        return mid
    } else if v <= ar[mid] {
        return binarySearch(v, ar, mid, end)
    } else {
        return binarySearch(v, ar, begin, mid-1)
    }
}

func calScore(x, y int, R []int) int {
    val := x*x + y*y
    if val > R[0] {
        return 0
    } else if val <= R[len(R)-1] {
        return len(R)
    }

    return binarySearch(val, R, 0, len(R)-1)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	var K, N int
	fmt.Fscan(reader, &K, &N)
    
    R := make([]int, K)
    for i:=0; i<K; i++ {
        fmt.Fscan(reader, &R[i])
    }
    for i:=0; i<K; i++ {
        R[i] *= R[i]
    }
    
    var x, y, score int
    for i:=0; i<N; i++ {
        fmt.Fscan(reader, &x, &y)
        score += calScore(x, y, R)
    }
	
    fmt.Fprintln(writer, score)
	writer.Flush()
}
