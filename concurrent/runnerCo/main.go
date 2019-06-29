package main

import (
	"errors"
	"log"
	"os"
	"os/signal"
	"time"
)

// Runner 在给定的超时时间内执行一组任务,
// 并且在操作系统发送中断信号时结束这些任务
type Runner struct {
	// interrupt 通道报告从操作系统
	// 发送的信号
	interrupt chan os.Signal
	// complete 通道报告处理任务已经完成
	compelete chan error
	// timeout 报告处理任务已经超时
	timeout	  <-chan time.Time
	// tasks 持有一组以索引顺序依次执行的
	// 函数
	tasks []func(int)
}

var ErrTimeout  = errors.New("received out")

var ErrInterrupt = errors.New("received interrupt")

func New(d time.Duration) *Runner  {
	return &Runner{
		interrupt:make(chan os.Signal, 1),
		compelete:make(chan error),
		timeout:time.After(d),
	}
}

func (r *Runner)Add(tasks ...func(int))  {
	r.tasks = append(r.tasks, tasks...)
}

func (r *Runner)Start() error  {
	signal.Notify(r.interrupt,os.Interrupt)
	go func() {
		r.compelete <- r.run()
	}()

	select {
	case err := <-r.compelete:
		return err
	case <-r.timeout:
		return  ErrTimeout
	}
}

func (r *Runner)run() error  {
	for id,task := range r.tasks {
		if r.gotInterrupt() {
			return ErrInterrupt
			task(id)
		}
	}
	return nil
}

func (r *Runner)gotInterrupt() bool  {
	select {
	case <-r.interrupt:
		signal.Stop(r.interrupt)
		return true
	default:
		return false
	}
}

const timeout  = 3*time.Second

func main()  {
	log.Println("Starting work.")
	r := New(timeout)
	r.Add(createTask(),createTask(),createTask())

	if err := r.Start(); err != nil {
		switch err {
		case ErrTimeout:
			log.Println("Terminating due to timeout")
			os.Exit(1)
		case ErrInterrupt:
			log.Println("Terminating due to interrupt")
			os.Exit(2)
		}
	}
	log.Println("Process ended")
}

// createTask 返回一个根据 id
// 休眠指定秒数的示例任务
func createTask() func(int) {
	return func(id int) {
		log.Printf("Processor - Task #%d.", id)
		time.Sleep(time.Duration(id) * time.Second)
	}
}