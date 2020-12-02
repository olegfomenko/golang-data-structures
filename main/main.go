// E-Olymp problem:
// https://www.e-olymp.com/uk/problems/687

package main

import (
	"../treap"
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	var treap treap.Treap

	for i := 0; i < n; i++ {
		var t, val int
		fmt.Fscan(in, &t, &val)

		switch t {
		case 1:
			treap.Insert(val, rand.Int())
		case -1:
			treap.Delete(val)
		case 0:
			fmt.Println(treap.Get(val))
		}
	}

}
