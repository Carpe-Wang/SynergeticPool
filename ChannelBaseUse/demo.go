package main

import "fmt"

func main() {
	var ch chan int
	fmt.Println(ch)      //<nil>
	fmt.Println(len(ch)) //0
	fmt.Println(cap(ch)) //0

	ch = make(chan int, 8)
	fmt.Println(ch)      //0xc00007c000
	fmt.Println(len(ch)) //0
	fmt.Println(cap(ch)) //8

	ch <- 8
	ch <- 10
	fmt.Println(ch)      //0xc00007c000
	fmt.Println(len(ch)) //2
	fmt.Println(cap(ch)) //8

	ch <- 8
	ch <- 10
	ch <- 8
	ch <- 10
	ch <- 8
	ch <- 10
	//ch <- 8 //在这就会报错，因为环形队列无法进行扩容
	//ch <- 10
	close(ch) //一定要先关闭chan
	for ele := range ch {
		fmt.Println(ele)
	}
}
