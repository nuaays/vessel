package models

import "github.com/jinzhu/gorm"

type Project struct {
	gorm.Model
	Name        string `json:"name"`
	Code        string `json:"code"`
	Actived     bool   `json:"actived"`
	SelfLink    string `json:"selfLink" gorm:"type:varchar(512)"`
	Created     int64  `json:"created"`
	Updated     int64  `json:"updated"`
	Description string `json:"description"`
}
