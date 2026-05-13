package userdetail

import (
	"go-hris/models/user"
	"time"
)

type View struct {
	Id          *int       `db:"id" json:"id" gorm:"primaryKey"`
	UserID      *int       `db:"user_id" json:"userID"`
	FullName    *string    `db:"full_name" json:"fullName"`
	Address     *string    `db:"address" json:"address"`
	JobPosition *string    `db:"job_position" json:"jobPosition"`
	Salary      *string    `db:"salary" json:"salary"`
	DateOfBirth *string    `db:"date_of_birth" json:"dateOfBirth"`
	JoinDate    *string    `db:"join_date" json:"joinDate"`
	LeaveAmount *int       `db:"leave_amount" json:"leaveAmount"`
	Note        *string    `db:"note" json:"note"`
	CreatedAt   *time.Time `db:"created_at" json:"createdAt"`

	UpdatedAt *time.Time `db:"updated_at" json:"updatedAt,omitempty"`

	UserDetail *user.View `db:"users" json:"user" gorm:"foreignKey:UserID;references:Id"`
}

type Create struct {
	UserID      *int    `db:"user_id" json:"userID" binding:"required"`
	FullName    *string `db:"full_name" json:"fullName" binding:"required"`
	Address     *string `db:"address" json:"address" binding:"required"`
	JobPosition *string `db:"job_position" json:"jobPosition" binding:"required"`
	Salary      *string `db:"salary" json:"salary" binding:"required"`
	DateOfBirth *string `db:"date_of_birth" json:"dateOfBirth" binding:"required"`
	JoinDate    *string `db:"join_date" json:"joinDate" binding:"required"`

	LeaveAmount *int      `db:"leave_amount" json:"leaveAmount"`
	Note        *string   `db:"note" json:"note"`
	CreatedAt   time.Time `db:"created_at" json:"createdAt"`
}

type Update struct {
	UserID      *int    `db:"user_id" json:"userID"`
	FullName    *string `db:"full_name" json:"fullName"`
	Address     *string `db:"address" json:"address"`
	JobPosition *string `db:"job_position" json:"jobPosition"`
	Salary      *string `db:"salary" json:"salary"`
	DateOfBirth *string `db:"date_of_birth" json:"dateOfBirth"`
	JoinDate    *string `db:"join_date" json:"joinDate"`
	LeaveAmount *int    `db:"leave_amount" json:"leaveAmount"`
	Note        *string `db:"note" json:"note"`

	UpdatedAt time.Time `db:"updated_at" json:"updatedAt"`
}

type Delete struct {
	Id int `db:"id" gorm:"primaryKey"`
}
