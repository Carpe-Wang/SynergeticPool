package main

import (
	"fmt"
	"time"
)

//-------------------------------有关Task的任务----------------------------//
//定一个任务类型
type Task struct {
	f func() error //一个task里面需要有一个具体的业务逻辑

}

//创建一个Task任务
func NewTask(arg_f func() error) *Task {
	t := Task{
		f: arg_f,
	}
	return &t
}

//Task也需要一个执行业务方法

func (t *Task) Execute() {
	t.f() //调用任务中已经绑定好的业务方法
}

//-------------------------------有关协程池的任务----------------------------//

type Pool struct {
	//对外的Task入口EntryChannel
	EntryChannel chan *Task

	//内部Task管道
	JobsChannel chan *Task

	//协程池中最大的协程数量
	Worker_num int
}

//创建pool函数

func NewPool(cap int) *Pool {
	p := Pool{
		EntryChannel: make(chan *Task),
		JobsChannel:  make(chan *Task),
		Worker_num:   cap,
	}
	return &p
}

//协程池创建Worker，并让这个worker去工作
func (p *Pool) worker(workerId int) {
	//一个worker的细节逻辑
	//从JobsChannel中取任务
	for task := range p.JobsChannel {
		//task就是当前worker的任务
		//得到任务，执行任务
		task.Execute()
	}
	fmt.Println("workerID为：", workerId, "执行完了一个任务")
}

//让协程池真正的开始工作，协程池一个方法
func (p *Pool) run() {
	//根据worker_num来创建worker去工作
	for i := 0; i < p.Worker_num; i++ {
		//每个worker都应该是一个goroutine
		go p.worker(i)
	}
	//从EntryChannel中取任务，将取到的任务，发给JobsChannel
	for task := range p.EntryChannel {
		p.JobsChannel <- task
	}
}

//用主函数来测试协程池的工作
func main() {
	//创建一些任务
	t := NewTask(func() error {
		//当前任务的业务，打印出当前的
		fmt.Println(time.Now())
		return nil
	})
	//创建一个Pool协程池
	p := NewPool(4)
	go func() {
		for {
			p.EntryChannel <- t
		}
	}()
	//将这些任务，交给协程池Pool
	p.run()
}
