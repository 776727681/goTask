package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var counter int64    // 原子操作的计数器
	var wg sync.WaitGroup

	// 启动10个协程
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// 每个协程执行1000次递增
			for j := 0; j < 1000; j++ {
				atomic.AddInt64(&counter, 1)
			}
		}()
	}

	wg.Wait() // 等待所有协程完成

	// 最终输出结果（原子读取）
	fmt.Printf("Final counter value: %d\n", atomic.LoadInt64(&counter))
}
