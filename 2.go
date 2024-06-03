package main

import "fmt"

func main() {
	var num, max, counter int
	for fmt.Scan(&num); num != 0; fmt.Scan(&num) {
		switch {
		case num > max:
			max = num
			counter = 1
		case num == max:
			counter++
		}
	}
	fmt.Print(counter)
}
