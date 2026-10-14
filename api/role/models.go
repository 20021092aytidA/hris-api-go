package role

import "time"

func (View) TableName() string {
	return "roles"
}

type Pagination struct {
	Page  int `form:"page"`
	Limit int `form:"limit"`
}

type AllowedParam struct {
	Id       string `form:"id"`
	RoleName string `form:"roleName"`
}

type AllParam struct {
	AllowedParam
	Pagination
}

type View struct {
	Id        *int       `db:"id" json:"id" gorm:"primaryKey"`
	RoleName  *string    `db:"role_name" json:"roleName"`
	CreatedAt *time.Time `db:"created_at" json:"createdAt"`

	UpdatedAt *time.Time `db:"updated_at" json:"updatedAt,omitempty"`
}

type Create struct {
	RoleName  *string   `db:"role_name" json:"roleName" binding:"required"`
	CreatedAt time.Time `db:"created_at" json:"createdAt" gorm:"autoCreateTime"`
}

type Update struct {
	RoleName  *string   `db:"role_name" json:"roleName"`
	UpdatedAt time.Time `db:"updated_at" json:"updatedAt" gorm:"autoUpdateTime"`
}
