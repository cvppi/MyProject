package main

import (
	"fmt"
)

func main() {
	var x, y, sum int
	fmt.Scan(&y)

	sli := []int{}

	for i := 0; i < y; i++ {
		fmt.Scan(&x)
		sli = append(sli, x)
	}

	myMin := sli[0]

	for i := range sli {
		if sli[0] < myMin {
			myMin = sli[i]
		}
	}
	// y = 1

	/*for i := 0; i < len(sli); i++ {
		if sli[y] < sli[i] || sli[y] == sli[i] {
			x = sli[y]
		}

		y++

		if y == len(sli) {
			y--
		}

	}*/

	for _, a := range sli {
		if myMin == a {
			sum++
		}
	}
	fmt.Println(sum)
}
