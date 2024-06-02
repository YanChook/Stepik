package main

import "fmt"

func main() {
	var n, b, sum int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&b)
		if 10 <= b && b%8 == 0 && b <= 99 {
			sum += b
		}
		fmt.Println(sum)

	}
}
