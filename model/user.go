package model

import "gorm.io/gorm"

type User struct {
	gorm.Model `json:"gorm_._model"`
	Username   string `json:"username,omitempty"`
	Password   string `json:"password,omitempty"`
	Phone      string `json:"phone,omitempty"`
}
