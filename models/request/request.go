package request

import (
	"go-hris/models/user"
	"time"
)

func (View) TableName() string {
	return "requests"
}

type Pagination struct {
	Page  int `form:"page"`
	Limit int `form:"limit"`
}

type AllowedParam struct {
	Id          int    `form:"id"`
	UserID      int    `form:"userID"`
	Title       string `form:"title"`
	Description string `form:"description"`
	Status      string `form:"status"`
}

type AllParam struct {
	AllowedParam
	Pagination
}

type View struct {
	Id          *int       `db:"id" json:"id" gorm:"primaryKey"`
	UserID      *int       `db:"user_id" json:"userID"`
	Title       *string    `db:"title" json:"title"`
	Description *string    `db:"description" json:"description"`
	Status      *string    `db:"status" json:"status"`
	CreatedAt   *time.Time `db:"created_at" json:"createdAt"`

	UpdatedAt *time.Time `db:"updated_at" json:"updatedAt,omitempty"`

	User *user.View `db:"users" json:"user" gorm:"foreignKey:UserID;references:Id"`
}

type Create struct {
	UserID      *int      `db:"user_id" json:"userID" binding:"required"`
	Title       *string   `db:"title" json:"title" binding:"required"`
	Description *string   `db:"description" json:"description" binding:"required"`
	CreatedAt   time.Time `db:"created_at" json:"createdAt" gorm:"autoCreateTime"`
}

type Update struct {
	Status    *string   `db:"status" json:"status" binding:"required"`
	UpdatedAt time.Time `db:"updated_at" json:"updatedAt" gorm:"autoUpdateTime"`
}

type Delete struct {
	Id int `db:"id" gorm:"primaryKey"`
}
