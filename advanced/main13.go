package main

import (
	"fmt"
	"sync"
	"time"
)

//任务结构体  包含任务ID、执行函数、耗时和错误信息
type Task struct {
	ID int
	Function func() error //任务函数
	Duration time.Duration //执行耗时
	Error   error //执行错误
}

//任务调度器  管理任务队列和并发执行
type TaskScheduler struct {
	tasks []*Task
	results chan *Task
	wg sync.WaitGroup // 使用 sync.WaitGroup 等待所有任务完成
}

//创建调度器
func NewScheduler() *TaskScheduler{
	return &TaskScheduler{
		results: make(chan *Task),
	}
}

//添加任务
func (s *TaskScheduler) AddTask(id int, f func() error)  {
	s.tasks = append(s.tasks, &Task{
		ID: id,
		Function: f,
	})
}

//执行任务
func (s *TaskScheduler) Run()  {
	//启动所有任务协程

	for _, task := range s.tasks {
		s.wg.Add(1)
		go func(t *Task) { //每个任务在独立 goroutine 中执行
			defer s.wg.Done()
			start := time.Now()
			err := t.Function()
			t.Duration = time.Since(start)
			t.Error = err
			s.results <- t
		}(task)
	}

	//等待所有任务完成
	go func() {
		s.wg.Wait()
		close(s.results)
	}()
}

// 获取结果   通过 channel (results) 安全收集结果
func (s *TaskScheduler) Results() <-chan *Task {
	return s.results
}

//添加任务 → 启动协程 → 记录时间 → 收集结果 → 关闭通道
func main()  {
	//创建调度器
	scheduler := NewScheduler()

	//添加实例任务
	scheduler.AddTask(1, func() error {
		time.Sleep(time.Second * 1)
		return nil
	})
	scheduler.AddTask(2, func() error {
		time.Sleep(time.Millisecond * 500)
		return fmt.Errorf("模拟错误")
	})
	scheduler.AddTask(3, func() error {
		time.Sleep(2 * time.Second)
		return nil
	})
	// 启动任务执行
	scheduler.Run()
	// 收集并显示结果
	for result := range scheduler.Results() {
		status := "成功"
		if result.Error != nil {
			status = "失败"
		}
		fmt.Printf("任务 %d: 状态 %s, 耗时 %v, 错误: %v\n",
			result.ID, status, result.Duration, result.Error)
	}
}
