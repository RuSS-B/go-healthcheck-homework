package model

import "github.com/jackc/pgtype"

func (User) TableName() string {
	return "users"
}

type User struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Username  pgtype.Varchar `json:"username" gorm:"unique"`
	FirstName *pgtype.Text   `json:"firstName"`
	LastName  *pgtype.Text   `json:"lastName"`
	Email     *pgtype.Text   `json:"email"`
	Phone     *pgtype.Text   `json:"phone"`
}
