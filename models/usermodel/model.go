package usermodel

import "time"

type ViewUser struct {
	UserID            *int       `db:"user_id" json:"userID"`
	RoleID            *int       `db:"role_id" json:"roleID"`
	FullName          *string    `db:"full_name" json:"fullName"`
	Address           *string    `db:"address" json:"address"`
	Position          *string    `db:"position" json:"position"`
	DateOfBirth       *string    `db:"date_of_birth" json:"dateOfBirth"`
	JoinDate          *string    `db:"join_date" json:"joinDate"`
	Salary            *int       `db:"salary" json:"salary"`
	LeaveAmount       *int       `db:"leave_amount" json:"leaveAmount"`
	Picture           *string    `db:"picture" json:"picture"`
	Resume            *string    `db:"resume" json:"resume"`
	DegreeCertificate *string    `db:"degree_certificate" json:"degreeCertificate"`
	ScoreTranscript   *string    `db:"score_transcript" json:"scoreTranscript"`
	Note              *string    `db:"note" json:"note"`
	CreatedAt         *time.Time `db:"created_at" json:"createdAt"`
	CreatedBy         *time.Time `db:"created_by" json:"createdBy"`
	EditedAt          *time.Time `db:"edited_at" json:"editedAt"`
	EditedBy          *time.Time `db:"edited_by" json:"editedBy"`
	DeletedAt         *time.Time `db:"deleted_at" json:"deletedAt"`
	DeletedBy         *time.Time `db:"deleted_by" json:"deletedBy"`
	IsDeleted         *string    `db:"is_deleted" json:"isDeleted"`
}

type CreateUser struct {
	RoleID            *int       `form:"role_id" db:"role_id" json:"roleID" binding:"required"`
	FullName          *string    `form:"full_name" db:"full_name" json:"fullName" binding:"required"`
	Address           *string    `form:"address" db:"address" json:"address" binding:"required"`
	Position          *string    `form:"position" db:"position" json:"position" binding:"required"`
	DateOfBirth       *string    `form:"date_of_birth" db:"date_of_birth" json:"dateOfBirth" binding:"required"`
	JoinDate          *string    `form:"join_date" db:"join_date" json:"joinDate" binding:"required"`
	Salary            *int       `form:"salary" db:"salary" json:"salary"`
	LeaveAmount       *int       `form:"leave_amount" db:"leave_amount" json:"leaveAmount"`
	Picture           *string    `form:"picture" db:"picture" json:"picture"`
	Resume            *string    `form:"resume" db:"resume" json:"resume"`
	DegreeCertificate *string    `form:"degree_certificate" db:"degree_certificate" json:"degreeCertificate"`
	ScoreTranscript   *string    `form:"score_transcript" db:"score_transcript" json:"scoreTranscript"`
	Note              *string    `form:"note" db:"note" json:"note"`
	CreatedAt         *time.Time `form:"created_at" db:"created_at" json:"createdAt" binding:"required"`
	CreatedBy         *time.Time `form:"created_by" db:"created_by" json:"createdBy" binding:"required"`
}

type UpdateUser struct {
	RoleID            *int       `form:"role_id" db:"role_id" json:"roleID"`
	FullName          *string    `form:"full_name" db:"full_name" json:"fullName"`
	Address           *string    `form:"address" db:"address" json:"address"`
	Position          *string    `form:"position" db:"position" json:"position"`
	DateOfBirth       *string    `form:"date_of_birth" db:"date_of_birth" json:"dateOfBirth"`
	JoinDate          *string    `form:"join_date" db:"join_date" json:"joinDate"`
	Salary            *int       `form:"salary" db:"salary" json:"salary"`
	LeaveAmount       *int       `form:"leave_amount" db:"leave_amount" json:"leaveAmount"`
	Picture           *string    `form:"picture" db:"picture" json:"picture"`
	Resume            *string    `form:"resume" db:"resume" json:"resume"`
	DegreeCertificate *string    `form:"degree_certificate" db:"degree_certificate" json:"degreeCertificate"`
	ScoreTranscript   *string    `form:"score_transcript" db:"score_transcript" json:"scoreTranscript"`
	Note              *string    `form:"note" db:"note" json:"note"`
	EditedAt          *time.Time `form:"edited_at" db:"edited_at" json:"editedAt" binding:"required"`
	EditedBy          *time.Time `form:"edited_by" db:"edited_by" json:"editedBy" binding:"required"`
}
