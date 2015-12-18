package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

var (
	zero    = big.NewInt(0)
	one     = big.NewInt(1)
	two     = big.NewInt(2)
	three   = big.NewInt(3)
	twelve  = big.NewInt(12)
	bignine = big.NewInt(1000000009)
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var T int
	fmt.Fscan(reader, &T)
	for i := 0; i < T; i++ {
		N := big.NewInt(0)
		K := big.NewInt(0)
		C := big.NewInt(0)
		fmt.Fscan(reader, N, K, C)
		N.Mod(N, bignine)
		K.Mod(K, bignine)
		C.Mod(C, bignine)
		temp := big.NewInt(1)
		ans := big.NewInt(1)
		temp.Mul(two, N)
		temp.Sub(temp, one)
		temp.Mul(temp, K)
		C.Mul(C, two)
		ans.Add(K, C)
		ans.Mul(ans, three)
		ans.Add(ans, temp)
		ans.Mul(ans, N)
		N.Sub(N, one)
		ans.Mul(ans, N)
		ans.Div(ans, twelve)
		ans.Mod(ans, bignine)
		fmt.Fprintln(writer, ans)
	}
}
