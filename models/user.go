package models

import "github.com/jinzhu/gorm"

type Organization struct {
	gorm.Model
	Name        string `json:"name"`
	Code        string `json:"code"`
	Email       string `json:"email"`
	Actived     bool   `json:"actived"`
	Description string `json:"description"`
	Teams       []Team
}

type Team struct {
	gorm.Model
	Name         string `json:"name"`
	Code         string `json:"code"`
	Email        string `json:"email"`
	Actived      bool   `json:"actived"`
	Description  string `json:"description"`
	Organization Organization
	Users        []User
}

type User struct {
	gorm.Model
	Name          string `json:"name"`
	Code          string `json:"code"`
	Email         string `json:"email"`
	Mobile        string `json:"mobile"`
	Actived       bool   `json:"actived"`
	Description   string `json:"description"`
	Teams         []Team
	Organizations []Organization
}

type Role struct {
	gorm.Model
	Name          string `json:"name"`
	Code          string `json:"code"`
	Actived       bool   `json:"actived"`
	Description   string `json:"description"`
	Teams         []Team
	Admin         bool
	Read          bool
	Write         bool
	Organizations []Organization
}
