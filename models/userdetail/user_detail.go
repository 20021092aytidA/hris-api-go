package userdetail

import "time"

type View struct {
	Id          *int       `db:"id" json:"id" gorm:"primaryKey"`
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
}

type Create struct {
	FullName    *string    `db:"full_name" json:"fullName" binding:"required"`
	Address     *string    `db:"address" json:"address" binding:"required"`
	JobPosition *string    `db:"job_position" json:"jobPosition" binding:"required"`
	Salary      *string    `db:"salary" json:"salary" binding:"required"`
	DateOfBirth *string    `db:"date_of_birth" json:"dateOfBirth" binding:"required"`
	JoinDate    *string    `db:"join_date" json:"joinDate" binding:"required"`
	CreatedAt   *time.Time `db:"created_at" json:"createdAt" binding:"required"`

	LeaveAmount *int    `db:"leave_amount" json:"leaveAmount"`
	Note        *string `db:"note" json:"note"`
}

type Update struct {
	FullName    *string `db:"full_name" json:"fullName"`
	Address     *string `db:"address" json:"address"`
	JobPosition *string `db:"job_position" json:"jobPosition"`
	Salary      *string `db:"salary" json:"salary"`
	DateOfBirth *string `db:"date_of_birth" json:"dateOfBirth"`
	JoinDate    *string `db:"join_date" json:"joinDate"`
	LeaveAmount *int    `db:"leave_amount" json:"leaveAmount"`
	Note        *string `db:"note" json:"note"`

	UpdatedAt *time.Time `db:"updated_at" json:"updatedAt" binding:"required"`
}
