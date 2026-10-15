package role

import "time"

func (ViewModel) TableName() string {
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

type ViewModel struct {
	Id        *int       `db:"id" json:"id" gorm:"primaryKey"`
	RoleName  *string    `db:"role_name" json:"roleName"`
	CreatedAt *time.Time `db:"created_at" json:"createdAt"`

	UpdatedAt *time.Time `db:"updated_at" json:"updatedAt,omitempty"`
}

type CreateModel struct {
	RoleName  *string   `db:"role_name" json:"roleName" binding:"required"`
	CreatedAt time.Time `db:"created_at" json:"createdAt" gorm:"autoCreateTime"`
}

type UpdateModel struct {
	RoleName  *string   `db:"role_name" json:"roleName"`
	UpdatedAt time.Time `db:"updated_at" json:"updatedAt" gorm:"autoUpdateTime"`
}
