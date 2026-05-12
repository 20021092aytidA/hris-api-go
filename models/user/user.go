package user

import (
	"go-hrs/models/role"
	"go-hrs/models/userdetail"
	"time"
)

type View struct {
	Id           *int       `db:"id" json:"id" gorm:"primaryKey"`
	UserDetailID *int       `db:"user_detail_id" json:"user_detail_id"`
	RoleID       *int       `db:"role_id" json:"roleID"`
	Username     *string    `db:"username" json:"username"`
	Email        *string    `db:"email" json:"email"`
	CreatedAt    *time.Time `db:"created_at" json:"createdAt"`

	UpdatedAt *time.Time `db:"updated_at" json:"updatedAt,omitempty"`

	UserDetail *userdetail.View `db:"user_details" json:"userDetail" gorm:"foreignKey:UserDetailID;references:Id"`
	Role       *role.View       `db:"roles" json:"role" gorm:"foreignKey:RoleID;references:Id"`
}

type Create struct {
	RoleID    *int      `db:"role_id" json:"roleID" binding:"required"`
	Password  *string   `db:"password" json:"password" binding:"required"`
	Username  *string   `db:"username" json:"username" binding:"required"`
	Email     *string   `db:"email" json:"email" binding:"required"`
	CreatedAt time.Time `db:"created_at" json:"createdAt" binding:"required"`
}

type Update struct {
	RoleID   *int    `db:"role_id" json:"roleID"`
	Password *string `db:"password" json:"password"`
	Username *string `db:"username" json:"username"`
	Email    *string `db:"email" json:"email"`

	UpdatedAt *time.Time `db:"updated_at" json:"updatedAt" binding:"required"`
}
