package main

// Now, we concatenate columns in increasing order of indices to get an array
//
// [1, 3, 2, 5, 4, 6]
//
// Given this final array, you have to count how many distinct initial matrices are possible. Return the required answer modulo 109+7.
//
// Two matrices are distinct if:
//
// 	Either their dimensions are different.
// 	Or if any of the corresponding row in two matrices are different.
//
// Example:
//
// If input array is [1, 3, 2, 4], distinct initial matrices possible are:
//
// [1 3 2 4]
// -----------------------
// [1 2]
// [3 4]
// -----------------------
// [1 4]
// [3 2]
// -----------------------
// [3 2]
// [1 4]
// -----------------------
// [3 4]
// [1 2]
// -----------------------
//
// that is, a total of 5 matrices.
//

const modNum = 1000000007

func cntMatrix(A []int) int {
	if len(A) < 2 {
		return len(A)
	}

	sortCheck := make([]int, len(A))
	sortCheck[0] = 1
	for i := 1; i < len(A); i++ {
		if A[i] >= A[i-1] {
			sortCheck[i] = sortCheck[i-1] + 1
		} else {
			sortCheck[i] = 1
		}
	}

	count := 0
	for i := 1; i <= len(A); i++ {
		res := possibleMatrices(A, sortCheck, i)
		count += res % modNum
	}

	return count % modNum
}

func possibleMatrices(A, sortCheck []int, l int) int {
	if len(A)%l != 0 {
		return 0
	}

	rows := l
	cols := len(A) / l
	for i := 0; i < cols; i++ {
		start := i * rows
		end := (i + 1) * rows
		if !sorted(sortCheck, start, end) {
			return 0
		}
	}

	return powmod((factmod(rows, modNum)), cols, modNum)
}

func powmod(x, y, p int) int {
	res := 1
	x = x % p
	if x == 0 {
		return 0
	}

	for y > 0 {
		// y is odd
		if y&1 != 0 {
			res = (res * x) % p
		}

		// y must be even now
		y = y >> 1
		x = (x * x) % p
	}

	return res
}

func factmod(n, p int) int {
	res := 1
	for n > 1 {
		var v int
		if (n/p)%2 == 0 {
			v = 1
		} else {
			v = p - 1
		}

		res = (res * v) % p
		for i := 2; i <= n%p; i++ {
			res = (res * i) % p
		}
		n /= p
	}
	return res % p
}

// end exclusive
func sorted(sortCheck []int, start, end int) bool {
	return (start == end) ||
		((sortCheck[end-1] - sortCheck[start]) == (end - 1 - start))
}
