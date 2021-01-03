package main

import (
	"log"
	"time"
)

func fact(n uint64) uint64 {
	if n == 0 || n == 1 {
		return 1
	}
	return n * fact(n-1)
}

func duration(x time.Time, appName string) {
	elapsed := time.Since(x)
	log.Printf("Execution Time: %d", elapsed)
}

func main() {
	defer duration(time.Now(), "factoriel")
	for i := 0; i <= 10; i++ {
		log.Printf("fact(%d) =  %d", i, fact(uint64(i)))
	}
}

// Output:
// 2021/01/03 00:55:28 fact(0) =  1
// 2021/01/03 00:55:28 fact(1) =  1
// 2021/01/03 00:55:28 fact(2) =  2
// 2021/01/03 00:55:28 fact(3) =  6
// 2021/01/03 00:55:28 fact(4) =  24
// 2021/01/03 00:55:28 fact(5) =  120
// 2021/01/03 00:55:28 fact(6) =  720
// 2021/01/03 00:55:28 fact(7) =  5040
// 2021/01/03 00:55:28 fact(8) =  40320
// 2021/01/03 00:55:28 fact(9) =  362880
// 2021/01/03 00:55:28 fact(10) =  3628800
// 2021/01/03 00:55:28 Execution Time: 155335
