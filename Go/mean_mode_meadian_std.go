package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var N, i, ssum, sum int
	var median float64
	fmt.Fscan(reader, &N)
	nums := make([]int, N)
	for i = 0; i < N; i++ {
		fmt.Fscan(reader, &(nums[i]))
		sum += nums[i]
		ssum += nums[i] * nums[i]
	}

	mean := float64(sum) / float64(N)
	fmt.Fprintf(writer, "%.1f\n", mean)
	sort.Ints(nums)
	if (N / 2 * 2) == N {
		median = float64(nums[N/2-1]+nums[N/2]) / 2.0
	} else {
		median = float64(nums[N/2])
	}
	fmt.Fprintf(writer, "%.1f\n", median)

	mfreq := 1
	mode := nums[0]
	cfreq := 1
	cnum := nums[0]
	for i = 1; i < N; i++ {
		if nums[i] == cnum {
			cfreq += 1
		} else {
			if (cfreq > mfreq) || (cfreq == mfreq && cnum <= mode) {
				mfreq = cfreq
				mode = cnum
			}

			cnum = nums[i]
			cfreq = 1
		}
	}
	if (cfreq > mfreq) || (cfreq == mfreq && cnum <= mode) {
		mfreq = cfreq
		mode = cnum
	}
	fmt.Fprintln(writer, mode)
	variance := float64(ssum)/float64(N) - mean*mean
	fmt.Fprintf(writer, "%.1f\n", math.Sqrt(variance))
}
