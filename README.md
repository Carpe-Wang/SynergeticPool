# Channel

#### 管道的底层实现原理

* 管道底层是一个环形队列(先进先出)，send(插入)和recv(取走)，从同一位置，沿同一方向顺序执行。
* sendx表示最后一次插入元素的位置，recvx代表最后一次取走的元素的位置。
* 不可以扩容

#### 管道的声明和初始化
```go
var ch chan int //管道的声明。
ch =make(chan int,8)//初始化，环形队列的容量为8
```
#### send和recv
```go
ch <-1//向管道里写入数据send
ch <-2
ch <-3
v:=<-ch//读数据(从管道中读取数据recv)
v =<-ch
```

#### 遍历管道
```
close(ch)//遍历前，需要关闭管道，防止有重新的插入
for ele:=range ch{
    fmt.Println(ele)
}
```

# 协程池设计思路
>为什么需要协程池？
* 当我们无限制的开创goroutine的时候，会浪费上下文切换资源，所以创建一个协程池限制goroutine在高并发情况下的个数。