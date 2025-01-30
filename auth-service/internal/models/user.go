package models

type User struct {
	ID         int64  `json:"id" gorm:"primary_key;type:uuid;default:gen_random_uuid()"`
	Email      string `json:"email" gorm:"unique;not null"`
	FirstName  string `json:"first_name" gorm:"not null"`
	LastName   string `json:"last_name" gorm:"not null"`
	MiddleName string `json:"middle_name" gorm:"not null"`
	Password   string `json:"password" gorm:"not null"`
}
