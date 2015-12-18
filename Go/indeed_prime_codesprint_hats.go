// greedy algorithm (incorrect)

package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

type Pair struct {
	val   float64
	index int
	smax  int
}

type PairHeap []Pair

func (h PairHeap) Len() int           { return len(h) }
func (h PairHeap) Less(i, j int) bool { return h[i].val > h[j].val }
func (h PairHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *PairHeap) Push(x interface{}) {
	*h = append(*h, x.(Pair))
}
func (h *PairHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var N, K int
	fmt.Fscan(reader, &N, &K)
	X := make([]int, N)
	S := make([]int, N)
	A := make([]int, N)
	B := make([]int, N)
	hp := &PairHeap{}
	for i := 0; i < N; i++ {
		fmt.Fscan(reader, &X[i], &S[i], &A[i], &B[i])
		pair := Pair{
			val:   float64(X[i]*S[i]) / float64(A[i]+(S[i]-1)*B[i]),
			index: i,
			smax:  S[i],
		}
		heap.Push(hp, pair)
	}

	Profit := 0
	for K > 0 && hp.Len() > 0 {
		p := heap.Pop(hp).(Pair)
		cur := (K-A[p.index])/B[p.index] + 1
		if p.smax <= cur && K >= A[p.index] {
			Profit += X[p.index] * p.smax
			K -= A[p.index] + (p.smax-1)*B[p.index]
		} else if cur > 0 && K >= A[p.index] {
			p.smax = cur
			p.val = float64(X[p.index]*cur) / float64(A[p.index]+(cur-1)*B[p.index])
			heap.Push(hp, p)
		}
	}

	fmt.Fprint(writer, Profit)
}
