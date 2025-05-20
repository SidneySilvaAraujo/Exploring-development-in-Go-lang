package models

type Person struct {
	Id    int    `json:"id" gorm:"primaryKey`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}
