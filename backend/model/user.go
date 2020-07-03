package model

type User struct {
	ID			uint	`gorm:"primary_key" json:"id"`
	Name		string	`json:"name"`
	Role		string	`json:"role"`
	Password	string 	`json:"password"`
}

func (User) TableName() string { return "users" }