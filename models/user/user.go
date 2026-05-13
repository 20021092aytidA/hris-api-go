package user

import (
	"go-hris/models/role"
	"time"
)

type View struct {
	Id        *int       `db:"id" json:"id" gorm:"primaryKey"`
	RoleID    *int       `db:"role_id" json:"roleID"`
	Username  *string    `db:"username" json:"username"`
	Email     *string    `db:"email" json:"email"`
	CreatedAt *time.Time `db:"created_at" json:"createdAt"`

	UpdatedAt *time.Time `db:"updated_at" json:"updatedAt,omitempty"`

	Role *role.View `db:"roles" json:"role" gorm:"foreignKey:RoleID;references:Id"`
}

type Create struct {
	RoleID   *int    `db:"role_id" json:"roleID" binding:"required"`
	Password *string `db:"password" json:"password" binding:"required"`
	Username *string `db:"username" json:"username" binding:"required"`
	Email    *string `db:"email" json:"email" binding:"required"`

	CreatedAt time.Time `db:"created_at" json:"createdAt"`
}

type Update struct {
	RoleID    *int       `db:"role_id" json:"roleID"`
	Password  *string    `db:"password" json:"password"`
	Username  *string    `db:"username" json:"username"`
	Email     *string    `db:"email" json:"email"`
	UpdatedAt *time.Time `db:"updated_at" json:"updatedAt"`
}

type Delete struct {
	Id int `db:"id" gorm:"primaryKey"`
}
