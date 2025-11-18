package main

import (
	"MetaNode/task2"
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// 测试 task2 包中的 AddTen 函数
	var num = 5
	task2.AddTen(&num)
	fmt.Println("AddTen 测试:", num) // 输出: 15

	// 测试 task2 包中的 DoubleValues 函数
	nums := []int{1, 2, 3, 4, 5}
	task2.DoubleValues(&nums)
	fmt.Println("DoubleValues 测试:", nums) // 输出: [2 4 6 8 10]

	// 测试 task2 包中的 PrintOddEven 函数
	fmt.Println("PrintOddEven 测试:")
	task2.PrintOddEven()

	// 测试 task2 包中的 TaskScheduler 函数
	fmt.Println("TaskScheduler 测试:")
	task2.TaskScheduler([]func(){
		func() {
			// 随机等待 1-3 毫秒
			time.Sleep(time.Duration(1+rand.Intn(3)) * time.Millisecond)
			fmt.Println("任务1")
		},
		func() {
			time.Sleep(time.Duration(1+rand.Intn(5)) * time.Millisecond)
			fmt.Println("任务2")
		},
	})

	// 测试 task2 包中的 Shape 接口
	fmt.Println("Shape 接口测试:")
	rect := &task2.Rectangle{Width: 3, Height: 4}
	circle := &task2.Circle{Radius: 5}
	fmt.Printf("矩形面积: %.2f, 周长: %.2f\n", rect.Area(), rect.Perimeter())
	fmt.Printf("圆面积: %.2f, 周长: %.2f\n", circle.Area(), circle.Perimeter())

	// 测试 task2 包中的 Person 和 Employee 结构体
	fmt.Println("Person 和 Employee 结构体测试:")
	person := &task2.Person{Name: "张三", Age: 30}
	employee := &task2.Employee{Person: *person, EmployeeID: 1001}
	employee.PrintInfo()

	// 测试 task2 包中的 ChannelCommunicate 函数
	fmt.Println("ChannelCommunicate 测试:")
	task2.ChannelCommunicate()

	// 测试 task2 包中的 BufferedChannel 函数
	fmt.Println("BufferedChannel 测试:")
	task2.BufferedChannel()

	// 测试 task2 包中的 MutexCounter 函数
	fmt.Println("MutexCounter 测试:")
	task2.MutexCounter()

	// 测试 task2 包中的 AtomicCounter 函数
	fmt.Println("AtomicCounter 测试:")
	task2.AtomicCounter()
}
