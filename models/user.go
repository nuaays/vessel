package models

import "github.com/jinzhu/gorm"

const (
	OrganizationTeamPermissions_Owner  = 100
	OrganizationTeamPermissions_Member = 200
	OrganizationTeamPermissions_Normal = 300

	TeamUserPermissions_Maintainer = 100
	TeamUserPermissions_Normal     = 300

	ProjectTeamPermissions_Admin = 100
	ProjectTeamPermissions_Write = 200
	ProjectTeamPermissions_Read  = 300
)

type Organization struct {
	gorm.Model
	Name          string `json:"name"`
	Code          string `json:"code"`
	Email         string `json:"email"`
	Actived       bool   `json:"actived"`
	Description   string `json:"description"`
	CreaterUserId int64
}

type Team struct {
	gorm.Model
	Name           string `json:"name"`
	Code           string `json:"code"`
	Email          string `json:"email"`
	Actived        bool   `json:"actived"`
	Description    string `json:"description"`
	OrganizationId int64
}

type User struct {
	gorm.Model
	Name        string `json:"name"`
	Code        string `json:"code"`
	Email       string `json:"email"`
	Actived     bool   `json:"actived"`
	Description string `json:"description"`
}

type OrganizationTeam struct {
	gorm.Model
	OrganizationId int64
	TeamId         int64
	Permissions    int64
}

type TeamUser struct {
	gorm.Model
	TeamId      int64
	UserId      int64
	Permissions int64
}

type ProjectTeam struct {
	gorm.Model
	TeamId      int64
	ProjectId   int64
	Permissions int64
}
