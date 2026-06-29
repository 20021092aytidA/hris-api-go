package user

import (
	"errors"
	"go-hris/models/role"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func (user *Create) BeforeCreate(tx *gorm.DB) error {
	if *user.Password != *user.ConfirmPassword {
		return errors.New("password not match confirm password")
	}

	bytePass := []byte(*user.Password)
	hashedBytePass, errHash := bcrypt.GenerateFromPassword(bytePass, bcrypt.DefaultCost)
	if errHash != nil {
		return errHash
	}

	*user.Password = string(hashedBytePass)
	return nil
}

func (user *Update) BeforeUpdate(tx *gorm.DB) error {
	if user.Password != nil {
		//EMPTY
		if *user.Password == "" {
			return errors.New("missing password")
		}

		//CONFIRM PASS
		if *user.Password != *user.ConfirmPassword {
			return errors.New("password not match with confirm password")
		}

		bytePass := []byte(*user.Password)
		hashedBytePass, errHash := bcrypt.GenerateFromPassword(bytePass, bcrypt.DefaultCost)
		if errHash != nil {
			return errHash
		}

		stringPass := string(hashedBytePass)
		user.Password = &stringPass
		return nil
	}

	return nil
}

func (View) TableName() string {
	return "users"
}

type Pagination struct {
	Page  int `form:"page"`
	Limit int `form:"limit"`
}

type AllowedParam struct {
	Id       string `form:"id"`
	RoleID   string `form:"roleID"`
	Username string `form:"username"`
	Email    string `form:"email"`
}

type AllParam struct {
	AllowedParam
	Pagination
}

type ViewWithPass struct {
	Id        *int       `db:"id" json:"id" gorm:"primaryKey"`
	RoleID    *int       `db:"role_id" json:"roleID"`
	Username  *string    `db:"username" json:"username"`
	Password  *string    `db:"password" json:"password"`
	Email     *string    `db:"email" json:"email"`
	CreatedAt *time.Time `db:"created_at" json:"createdAt"`
	UpdatedAt *time.Time `db:"updated_at" json:"updatedAt,omitempty"`
	Role      *role.View `db:"roles" json:"role" gorm:"foreignKey:RoleID;references:Id"`
}

type View struct {
	Id        *int       `db:"id" json:"id" gorm:"primaryKey"`
	RoleID    *int       `db:"role_id" json:"roleID"`
	Username  *string    `db:"username" json:"username"`
	Email     *string    `db:"email" json:"email"`
	CreatedAt *time.Time `db:"created_at" json:"createdAt"`
	UpdatedAt *time.Time `db:"updated_at" json:"updatedAt,omitempty"`
	Role      *role.View `db:"roles" json:"role" gorm:"foreignKey:RoleID;references:Id"`
}

type Create struct {
	Id              uint      `db:"id" json:"id" gorm:"primaryKey;autoIncrement;column:id"`
	RoleID          *int      `db:"role_id" json:"roleID" binding:"required"`
	Password        *string   `db:"password" json:"password" binding:"required"`
	ConfirmPassword *string   `json:"confirmPassword" binding:"required"`
	Username        *string   `db:"username" json:"username" binding:"required"`
	Email           *string   `db:"email" json:"email" binding:"required"`
	CreatedAt       time.Time `db:"created_at" json:"createdAt" gorm:"autoCreateTime"`
}

type Update struct {
	RoleID          *int      `db:"role_id" json:"roleID"`
	Password        *string   `db:"password" json:"password"`
	ConfirmPassword *string   `json:"confirmPassword"`
	Username        *string   `db:"username" json:"username"`
	Email           *string   `db:"email" json:"email"`
	UpdatedAt       time.Time `db:"updated_at" json:"updatedAt" gorm:"autoUpdateTime"`
}

type Login struct {
	Username *string `db:"username" json:"username" binding:"required"`
	Password *string `db:"password" json:"password" binding:"required"`
}
