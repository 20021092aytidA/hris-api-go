package adminmodel

import "time"

type ViewAdmin struct {
	AdminID   *int       `db:"admin_id" json:"adminID"`
	UserID    *int       `db:"user_id" json:"userID"`
	RoleID    *int       `db:"role_id" json:"roleID"`
	Username  *string    `db:"username" json:"username"`
	Password  *string    `db:"password" json:"password"`
	CreatedAt *time.Time `db:"created_at" json:"createdAt"`
	CreatedBy *int       `db:"created_by" json:"createdBy"`
	EditedAt  *time.Time `db:"edited_at" json:"editedAt"`
	EditedBy  *int       `db:"edited_by" json:"editedBy"`
	DeletedAt *time.Time `db:"deleted_at" json:"deletedAt"`
	DeletedBy *int       `db:"deleted_by" json:"deletedBy"`
	IsDeleted *string    `db:"is_deleted" json:"isDeleted"`
}

type CreateAdmin struct {
	UserID    *int       `form:"user_id" db:"user_id" json:"userID" binding:"required"`
	RoleID    *int       `form:"role_id" db:"role_id" json:"roleID" binding:"required"`
	Username  *string    `form:"username" db:"username" json:"username" binding:"required"`
	Password  *string    `form:"password" db:"password" json:"password" binding:"required"`
	CreatedAt *time.Time `form:"created_at" db:"created_at" json:"createdAt" binding:"required" time_format:"2006-01-02 15:04:05"`
	CreatedBy *int       `form:"created_by" db:"created_by" json:"createdBy" binding:"required"`
}

type LoginAdmin struct {
	Username *string `form:"username" db:"username" json:"username" binding:"required"`
	Password *string `form:"password" db:"password" json:"password" binding:"required"`
}

type UpdateAdmin struct {
	UserID   *int       `form:"user_id" db:"user_id" json:"userID"`
	RoleID   *int       `form:"role_id" db:"role_id" json:"roleID"`
	Username *string    `form:"username" db:"username" json:"username"`
	Password *string    `form:"password" db:"password" json:"password"`
	EditedAt *time.Time `form:"edited_at" db:"edited_at" json:"editedAt" binding:"required" time_format:"2006-01-02 15:04:05"`
	EditedBy *int       `form:"edited_by" db:"edited_by" json:"editedBy" binding:"required"`
}
