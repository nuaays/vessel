package models

import "github.com/jinzhu/gorm"

type Workspace struct {
	gorm.Model
	Name        string `json:"name"`
	Actived     bool   `json:"actived"`
	Description string `json:"description"`
	projects    []Project
}
