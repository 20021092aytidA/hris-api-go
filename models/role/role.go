package role

import "time"

type View struct {
	Id        *int       `db:"id" json:"id" gorm:"primaryKey"`
	RoleName  *string    `db:"role_name" json:"roleName"`
	CreatedAt *time.Time `db:"created_at" json:"createdAt"`

	UpdatedAt *time.Time `db:"updated_at" json:"updatedAt,omitempty"`
}

type Create struct {
	RoleName  *string    `db:"role_name" json:"roleName" binding:"required"`
	CreatedAt *time.Time `db:"created_at" json:"createdAt" binding:"required"`
}

type Update struct {
	RoleName *string `db:"role_name" json:"roleName"`

	UpdatedAt *time.Time `db:"updated_at" json:"updatedAt" binding:"required"`
}
