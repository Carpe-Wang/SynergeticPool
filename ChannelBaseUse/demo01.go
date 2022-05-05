package main

import "fmt"

func main() {

	ch := make(chan int, 10)
	ch <- 10
	go student(ch)
	for {

	}
}

func student(ch chan int) {
	fmt.Println("我是一个学生", <-ch)
}
