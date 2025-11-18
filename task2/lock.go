package task2

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
func MutexCounter() {
	// 创建一个共享的计数器
	var counter int
	var locker sync.Mutex

	// 启动10个协程，每个协程对计数器进行1000次递增操作
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				locker.Lock()
				counter++
				locker.Unlock()
			}
		}()
	}

	// 等待所有协程完成
	wg.Wait()

	// 输出最终的计数器值
	fmt.Println("Counter Value:", counter)
}

// 使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
func AtomicCounter() {
	// 创建一个共享的原子计数器
	var counter int64

	// 启动10个协程，每个协程对计数器进行1000次递增操作
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				atomic.AddInt64(&counter, 1)
			}
		}()
	}

	// 等待所有协程完成
	wg.Wait()

	// 输出最终的计数器值
	fmt.Println("Counter Value:", atomic.LoadInt64(&counter))
}
