package model

import (
	"github.com/guregu/null"
)

func (User) TableName() string {
	return "users"
}

type User struct {
	ID        uint        `json:"id" gorm:"primaryKey"`
	Username  null.String `json:"username" gorm:"unique"`
	FirstName null.String `json:"firstName"`
	LastName  null.String `json:"lastName"`
	Email     null.String `json:"email"`
	Phone     null.String `json:"phone"`
}
