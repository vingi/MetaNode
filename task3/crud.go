package task3

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 假设有一个名为 students 的表，包含字段 id （主键，自增）、 name （学生姓名，字符串类型）、 age （学生年龄，整数类型）、 grade （学生年级，字符串类型）。
type Student struct {
	ID    uint `gorm:"primaryKey"`
	Name  string
	Age   int
	Grade string
}

// 创建 students 表
func CreateTable() {
	db := ConnectDB()
	// 创建 students 表
	db.AutoMigrate(&Student{})
}

// 连接数据库
func ConnectDB() *gorm.DB {
	// 连接数据库
	dsn := "root:520530@tcp(192.168.1.2:3306)/metanode?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败")
	}
	return db
}

// 要求 ：
// 编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
func InsertStudent() {
	// 假设使用的是 MySQL 数据库
	db := ConnectDB()
	// query := "INSERT INTO students (name, age, grade) VALUES ('张三', 20, '三年级')"
	student := Student{Name: "张三", Age: 20, Grade: "三年级"}
	result := db.Create(&student)
	fmt.Println(result.RowsAffected)
}

// 编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
func QueryStudents() {
	// 假设使用的是 MySQL 数据库
	db := ConnectDB()
	var students []Student
	db.Where("age > ?", 18).Find(&students)
	fmt.Println(students)
	// query := "SELECT * FROM students WHERE age > 18"

}

// 编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
func UpdateStudentGrade() {
	// 假设使用的是 MySQL 数据库
	db := ConnectDB()
	var student Student
	db.Where("name = ?", "张三").First(&student)
	student.Grade = "四年级"
	db.Save(&student)
	// query := "UPDATE students SET grade = '四年级' WHERE name = '张三'"

}

// 编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。
func DeleteStudents() {
	// 假设使用的是 MySQL 数据库
	db := ConnectDB()
	db.Where("age < ?", 15).Delete(&Student{})
	// query := "DELETE FROM students WHERE age < 15"

}
