package main

import (
	"time"
	"fmt"
)

func say(s string){
	for i:= 0; i < 5; i++{
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func sum(s []int, c chan int){
	sum := 0
	for _, v := range s{
		sum += v
	}
	c <- sum
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main(){
	go say("World")
	say("Hello")

	s := []int{5, 8, 3, 6, -1, 0, 4}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <- c, <-c

	fmt.Println(x, y, x+ y)

	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
//	ch <- 3
	fmt.Println(<-ch)
	fmt.Println(<-ch)
//	fmt.Println(<-ch)

	c = make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}
