package task3

import (
	"fmt"

	"gorm.io/gorm"
)

// 题目1：模型定义
// 假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
// 要求 ：
// 使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章）， Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
// 定义 User 模型
type User struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Email     string `gorm:"unique"`
	Password  string
	PostCount int
}

// 定义 Post 模型
type Post struct {
	ID       uint `gorm:"primaryKey"`
	Title    string
	Content  string
	AuthorID uint
	Author   User `gorm:"foreignKey:AuthorID;references:ID"`
	Comments []Comment
}

// 定义 Comment 模型
type Comment struct {
	ID      uint `gorm:"primaryKey"`
	Content string
	PostID  uint
	Post    Post `gorm:"foreignKey:PostID;references:ID"`
}

// 编写Go代码，使用Gorm创建这些模型对应的数据库表。
func CreateTables() error {
	db := ConnectDB()
	err := db.AutoMigrate(&User{}, &Post{}, &Comment{})
	if err != nil {
		fmt.Println("数据库表创建失败", err)
		return err
	}
	return nil
}

// 题目2：关联查询
// 基于上述博客系统的模型定义。
// 要求 ：
// 编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
func QueryPostsByUserID(userID uint) ([]Post, error) {
	var posts []Post
	db := ConnectDB()
	err := db.Where("author_id = ?", userID).Preload("Comments").Find(&posts).Error
	if err != nil {
		return posts, err
	}
	return posts, nil
}

// 编写Go代码，使用Gorm查询评论数量最多的文章信息。
func QueryPostWithMostComments() (Post, error) {
	var post Post
	db := ConnectDB()
	// 获取评论数量最多的文章
	err := db.Preload("Comments").Order("Comments DESC").First(&post).Error
	if err != nil {
		return post, err
	}
	return post, nil
}

// 题目3：钩子函数
// 继续使用博客系统的模型。
// 要求 ：
// 为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
func (p *Post) BeforeCreate(tx *gorm.DB) error {
	db := ConnectDB()
	// 更新用户的文章数量统计字段
	var user User
	db.Where("id = ?", p.AuthorID).First(&user)
	user.PostCount++
	db.Save(&user)
	return nil
}

// 为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
func (c *Comment) BeforeDelete(tx *gorm.DB) error {
	db := ConnectDB()
	// 检查文章的评论数量
	var post Post
	db.Where("id = ?", c.PostID).First(&post)
	if len(post.Comments) == 0 {
		post.Comments = []Comment{}
		db.Save(&post)
	}
	return nil
}
