package rolemodel

import "time"

func (ViewRole) TableName() string {
	return "role"
}

type ViewRole struct {
	RoleID    *int       `db:"role_id" json:"roleID" gorm:"primarykey"`
	RoleName  *string    `db:"role_name" json:"roleName"`
	CreatedAt *time.Time `db:"created_at" json:"createdAt"`
	CreatedBy *int       `db:"created_by" json:"createdBy"`
	EditedAt  *time.Time `db:"edited_at" json:"editedAt"`
	EditedBy  *int       `db:"edited_by" json:"editedBy"`
	DeletedAt *time.Time `db:"deleted_at" json:"deletedAt"`
	DeletedBy *int       `db:"deleted_by" json:"deletedBy"`
	IsDeleted *string    `db:"is_deleted" json:"isDeleted"`
}

type CreateRole struct {
	RoleName  *string    `form:"role_name"  binding:"required"`
	CreatedAt *time.Time `form:"created_at" db:"created_at" binding:"required" time_format:"2006-01-02 15:04:05"`
	CreatedBy *int       `form:"created_by" db:"created_by" binding:"required"`
}

type UpdateRole struct {
	RoleName *string    `form:"role_name" db:"role_name" json:"roleName"`
	EditedAt *time.Time `form:"edited_at" db:"edited_at" json:"editedAt" binding:"required" time_format:"2006-01-02 15:04:05"`
	EditedBy *int       `form:"edited_by" db:"edited_by" json:"editedBy" binding:"required"`
}
