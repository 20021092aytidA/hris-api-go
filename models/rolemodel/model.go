package rolemodel

import "time"

type ViewRole struct {
	RoleID    *int       `db:"role_id" json:"roleID"`
	RoleName  *string    `db:"role_name" json:"roleName"`
	CreatedAt *time.Time `db:"created_at" json:"createdAt"`
	CreatedBy *time.Time `db:"created_by" json:"createdBy"`
	EditedAt  *time.Time `db:"edited_at" json:"editedAt"`
	EditedBy  *time.Time `db:"edited_by" json:"editedBy"`
	DeletedAt *time.Time `db:"deleted_at" json:"deletedAt"`
	DeletedBy *time.Time `db:"deleted_by" json:"deletedBy"`
	IsDeleted *string    `db:"is_deleted" json:"isDeleted"`
}

type CreateRole struct {
	RoleName  *string    `form:"role_name" db:"role_name" json:"roleName" binding:"required"`
	CreatedAt *time.Time `form:"created_at" db:"created_at" json:"createdAt" binding:"required"`
	CreatedBy *time.Time `form:"created_by" db:"created_by" json:"createdBy" binding:"required"`
}

type UpdateRole struct {
	RoleName *string    `form:"role_name" db:"role_name" json:"roleName"`
	EditedAt *time.Time `form:"edited_at" db:"edited_at" json:"editedAt" binding:"required"`
	EditedBy *time.Time `form:"edited_by" db:"edited_by" json:"editedBy" binding:"required"`
}
