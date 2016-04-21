package models

import "github.com/jinzhu/gorm"

type Project struct {
	gorm.Model
	Created     int64  `json:"created"`
	Updated     int64  `json:"updated"`
	Name        string `json:"name"`
	Code        string `json:"code"`
	Actived     bool   `json:"actived"`
	Description string `json:"description"`
}
