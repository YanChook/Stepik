package main

import "fmt"

func main() {

	var x, y, p int

	fmt.Scan(&x, &p, &y)

	for i := 0; i < y; i++ {
		if x > y {
			fmt.Println(i)
			break
		} else {
			x = x + (x * p / 100)
		}
	}
}
