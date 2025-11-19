package task3

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// 连接数据库
func ConnectDB_sqlx() *sqlx.DB {
	dsn := "root:520530@tcp(192.168.1.2:3306)/metanode?parseTime=true"
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		panic("数据库连接失败")
	}
	return db
}

type Employee struct {
	ID         int     `db:"id"`
	Name       string  `db:"name"`
	Department string  `db:"department"`
	Salary     float64 `db:"salary"`
}

// 假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。
// 要求 ：
// 编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
func QueryEmployeesByDepartment(department string) ([]Employee, error) {
	var employees []Employee
	db := ConnectDB_sqlx()
	err := db.Select(&employees, "SELECT * FROM employees WHERE department = ?", department)
	if err != nil {
		return employees, err
	}
	return employees, nil
}

// 编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。
func QueryHighestSalaryEmployee() (Employee, error) {
	var employee Employee
	db := ConnectDB_sqlx()
	err := db.Get(&employee, "SELECT * FROM employees ORDER BY salary DESC LIMIT 1")
	if err != nil {
		return employee, err
	}
	return employee, nil
}

// 题目2：实现类型安全映射
// 假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。
// 定义一个 Book 结构体，包含与 books 表对应的字段。
type Book struct {
	ID     int     `db:"id"`
	Title  string  `db:"title"`
	Author string  `db:"author"`
	Price  float64 `db:"price"`
}

// 编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。
func QueryBooksByPrice(price float64) ([]Book, error) {
	var books []Book
	db := ConnectDB_sqlx()
	err := db.Select(&books, "SELECT * FROM books WHERE price > ?", price)
	if err != nil {
		return books, err
	}
	return books, nil
}
