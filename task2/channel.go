package task2

import "fmt"

//编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
func ChannelCommunicate() {
	// 创建一个整数通道
	ch := make(chan int)
	done := make(chan bool, 2)

	// 启动一个协程，生成从1到10的整数并发送到通道中
	go func() {
		for i := 1; i <= 10; i++ {
			ch <- i
		}
		close(ch) // 发送完所有整数后关闭通道
		done <- true
	}()

	// 启动另一个协程，从通道中接收整数并打印出来
	go func() {
		for num := range ch {
			fmt.Println(num)
		}
		done <- true
	}()

	// 等待两个协程完成
	<-done
	<-done
}

// 实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
func BufferedChannel() {
	// 创建一个带有缓冲的整数通道，缓冲区为10
	ch := make(chan int, 10)
	done := make(chan bool, 2)

	// 启动一个协程，生成从1到100的整数并发送到通道中
	go func() {
		for i := 1; i <= 100; i++ {
			ch <- i
		}
		close(ch) // 发送完所有整数后关闭通道
		done <- true
	}()

	// 启动另一个协程，从通道中接收整数并打印出来
	go func() {
		for num := range ch {
			fmt.Println(num)
		}
		done <- true
	}()

	// 等待两个协程完成
	<-done
	<-done
}
