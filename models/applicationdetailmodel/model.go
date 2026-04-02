package applicationdetailmodel

import "time"

type ViewApplicationDetail struct {
	ApplicantDetailID *int       `db:"applicant_detail_id" json:"applicantDetailID"`
	UserID            *int       `db:"user_id" json:"userID"`
	ExpectedSalary    *int       `db:"expected_salary" json:"expectedSalary"`
	PrevSalary        *int       `db:"prev_salary" json:"prevSalary"`
	PrevPosition      *string    `db:"prev_position" json:"prevPosition"`
	ApplyDate         *string    `db:"apply_date" json:"applyDate"`
	IsRejected        *string    `db:"is_rejected" json:"isRejected"`
	RejectedDate      *string    `db:"rejected_date" json:"rejectedDate"`
	Note              *string    `db:"note" json:"note"`
	CreatedAt         *time.Time `db:"created_at" json:"createdAt"`
	CreatedBy         *time.Time `db:"created_by" json:"createdBy"`
	EditedAt          *time.Time `db:"edited_at" json:"editedAt"`
	EditedBy          *time.Time `db:"edited_by" json:"editedBy"`
	DeletedAt         *time.Time `db:"deleted_at" json:"deletedAt"`
	DeletedBy         *time.Time `db:"deleted_by" json:"deletedBy"`
	IsDeleted         *string    `db:"is_deleted" json:"isDeleted"`
}

type CreateApplicationDetail struct {
	UserID         *int       `form:"user_id" db:"user_id" json:"userID" binding:"required"`
	ExpectedSalary *int       `form:"expected_salary" db:"expected_salary" json:"expectedSalary" binding:"required"`
	PrevSalary     *int       `form:"prev_salary" db:"prev_salary" json:"prevSalary"`
	PrevPosition   *string    `form:"prev_position" db:"prev_position" json:"prevPosition"`
	ApplyDate      *string    `form:"apply_date" db:"apply_date" json:"applyDate" binding:"required"`
	IsRejected     *string    `form:"is_rejected" db:"is_rejected" json:"isRejected"`
	RejectedDate   *string    `form:"rejected_date" db:"rejected_date" json:"rejectedDate"`
	Note           *string    `form:"note" db:"note" json:"note"`
	CreatedAt      *time.Time `form:"created_at" db:"created_at" json:"createdAt" binding:"required"`
	CreatedBy      *time.Time `form:"created_by" db:"created_by" json:"createdBy" binding:"required"`
}

type UpdateApplicationDetail struct {
	UserID         *int       `form:"user_id" db:"user_id" json:"userID"`
	ExpectedSalary *int       `form:"expected_salary" db:"expected_salary" json:"expectedSalary"`
	PrevSalary     *int       `form:"prev_salary" db:"prev_salary" json:"prevSalary"`
	PrevPosition   *string    `form:"prev_position" db:"prev_position" json:"prevPosition"`
	ApplyDate      *string    `form:"apply_date" db:"apply_date" json:"applyDate"`
	IsRejected     *string    `form:"is_rejected" db:"is_rejected" json:"isRejected"`
	RejectedDate   *string    `form:"rejected_date" db:"rejected_date" json:"rejectedDate"`
	Note           *string    `form:"note" db:"note" json:"note"`
	EditedAt       *time.Time `form:"edited_at" db:"edited_at" json:"editedAt" binding:"required"`
	EditedBy       *time.Time `form:"edited_by" db:"edited_by" json:"editedBy" binding:"required"`
}
