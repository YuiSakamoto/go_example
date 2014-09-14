package main

import "fmt"
import "time"

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum // send sum to c
}

func sum2(a []int, c chan int, w int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	time.Sleep(100 * time.Millisecond)
	c <- sum // send sum to c
}

func main() {
	a := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum2(a[:len(a)/2], c, 100)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}
