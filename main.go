package main

import "fmt"

func main() {
	ch := make(chan int)
	ch1 := make(chan string)
	go sayHello(ch)
	go Hello(ch1)

	res := <-ch
	fmt.Println(res)
	res1 := <-ch1
	fmt.Println(res1)

}

func sayHello(ch chan int) {
	ch <- 10

}

func Hello(ch chan string) {

	ch <- "Hii From Hello Goroutine" 

}
