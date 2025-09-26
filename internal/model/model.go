package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string    `gorm:"unique;not null"`
	Password string    `gorm:"not null"`
	Email    string    `gorm:"unique;not null"`
	Posts    []Post    `json:"-" gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Comments []Comment `json:"-" gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

type Post struct {
	gorm.Model
	Title    string `gorm:"not null"`
	Content  string `gorm:"not null"`
	UserID   uint
	User     User      `gorm:"foreignKey:UserID;references:ID"`
	Comments []Comment `gorm:"foreignKey:PostID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type Comment struct {
	gorm.Model
	Content string `gorm:"not null"`
	UserID  uint
	User    User `gorm:"foreignKey:UserID;references:ID"`
	PostID  uint
	Post    Post `gorm:"foreignKey:PostID;references:ID"`
}
