package task2

import (
	"fmt"
	"sync"
	"time"
)

// 编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
func PrintOddEven() {
	// 使用通道来等待协程完成
	// done := make(chan bool, 2)
	var wg sync.WaitGroup
	wg.Add(2)

	// 启动一个协程打印奇数
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i += 2 {
			fmt.Println(i)
		}
		// done <- true
	}()

	// 启动一个协程打印偶数
	go func() {
		defer wg.Done()
		for i := 2; i <= 10; i += 2 {
			fmt.Println(i)
		}
		// done <- true
	}()

	wg.Wait()
	// 等待两个协程都完成
	// <-done
	// <-done
}

// 设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
func TaskScheduler(tasks []func()) {
	// 使用通道来等待所有协程完成
	done := make(chan bool, len(tasks))

	for _, task := range tasks {
		go func(t func()) {
			start := time.Now()
			t()
			fmt.Printf("当前任务花费了 %v\n", time.Since(start))
			done <- true
		}(task)
	}

	// 等待所有任务完成
	for i := 0; i < len(tasks); i++ {
		<-done
	}
}
