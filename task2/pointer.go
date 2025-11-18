package task2

// 定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10
func AddTen(p *int) {
	*p += 10
}

// 实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
func DoubleValues(p *[]int) {
	for i := 0; i < len(*p); i++ {
		(*p)[i] *= 2
	}
}
