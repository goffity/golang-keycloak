package v1

import "gorm.io/gorm"

type Member struct {
	gorm.Model
	ID        int    `json:"ID"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       int    `json:"age"`
	Address   string `json:"address"`
}
