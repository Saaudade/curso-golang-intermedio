package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	go doSomething(c)
	<-c

	// Apuntadores
	g := 25
	fmt.Println(g)
	h := &g
	fmt.Println(h)
	fmt.Println(*h)
}

func doSomething(c chan int) {
	time.Sleep(3 * time.Second)
	fmt.Println("Done")
	c <- 1
}
