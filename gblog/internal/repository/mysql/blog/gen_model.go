package blog

import "time"

// Blog
//
//go:generate gormgen -structs Blog -input .
type Blog struct {
	AID        int32     //
	Author     string    //
	Category   string    //
	Title      string    //
	Content    string    //
	Status     int32     //
	AddTime    time.Time `gorm:"time"` // getdate()
	UpdateTime time.Time `gorm:"time"` // getdate()
}
