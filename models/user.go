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

	PipelineTeamPermissions_Admin = 100
	PipelineTeamPermissions_Write = 200
	PipelineTeamPermissions_Read  = 300
	PipelineTeamPermissions_Run   = 400
)

type Organization struct {
	gorm.Model
	Created       int64  `json:"created"`
	Updated       int64  `json:"updated"`
	Name          string `json:"name"`
	Code          string `json:"code"`
	Email         string `json:"email"`
	Actived       bool   `json:"actived"`
	Description   string `json:"description"`
	CreaterUserId int64
}

type Team struct {
	gorm.Model
	Created        int64  `json:"created"`
	Updated        int64  `json:"updated"`
	Name           string `json:"name"`
	Code           string `json:"code"`
	Email          string `json:"email"`
	Actived        bool   `json:"actived"`
	Description    string `json:"description"`
	OrganizationId int64
}

type User struct {
	gorm.Model
	Created        int64  `json:"created"`
	Updated        int64  `json:"updated"`
	Name           string `json:"name"`
	Code           string `json:"code"`
	Password       string `json:"password"`
	Token          string `json:"token"`
	TokenExpiresIn int64  `json:"tokenExpiresIn"`
	Email          string `json:"email"`
	Actived        bool   `json:"actived"`
	Description    string `json:"description"`
}

type OrganizationTeam struct {
	gorm.Model
	OrganizationId int64
	TeamId         int64
	Permissions    int64
}

type TeamUser struct {
	gorm.Model
	Created     int64 `json:"created"`
	Updated     int64 `json:"updated"`
	TeamId      int64
	UserId      int64
	Permissions int64
}

type ProjectTeam struct {
	gorm.Model
	Created     int64 `json:"created"`
	Updated     int64 `json:"updated"`
	TeamId      int64
	ProjectId   int64
	Permissions int64
}

type PipelineTeam struct {
	gorm.Model
	Created     int64 `json:"created"`
	Updated     int64 `json:"updated"`
	TeamId      int64
	PipelineId  int64
	Permissions int64
}
