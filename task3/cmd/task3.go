package main

import (
	"MetaNode/task3"
)

func main() {
	// CRUD
	// task3.CreateTable()
	// task3.InsertStudent()
	// task3.QueryStudents()
	// task3.UpdateStudentGrade()
	// task3.DeleteStudents()

	// 转账
	// task3.CreateTable2()
	// task3.InsertAccount()
	// task3.Transfer(1, 2, 100)

	// gormadvanced
	err := task3.CreateTables()
	if err != nil {
		panic("数据库表创建失败")
	}
}
