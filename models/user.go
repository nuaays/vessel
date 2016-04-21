package models

import "github.com/jinzhu/gorm"

const (
	OrganizationPermissions_Normal = 100
	OrganizationPermissions_Member = 200
	OrganizationPermissions_Owner  = 300

	TeamUserPermissions_Normal     = 100

	ProjectTeamPermissions_Read  = 100
	ProjectTeamPermissions_Write = 200

	PipelineTeamPermissions_Read  = 100
	PipelineTeamPermissions_Write = 200
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
	// Permissions:OrganizationPermissions_Normal = 100
	// Permissions:OrganizationPermissions_Member = 200
	// Permissions:OrganizationPermissions_Owner  = 300
	Permissions    int64
}

type TeamUser struct {
	gorm.Model
	Created     int64 `json:"created"`
	Updated     int64 `json:"updated"`
	TeamId      int64
	UserId      int64
	// Permissions:TeamUserPermissions_Normal     = 100
	Permissions int64
}

type ProjectTeam struct {
	gorm.Model
	Created     int64 `json:"created"`
	Updated     int64 `json:"updated"`
	TeamId      int64
	ProjectId   int64
	// Permissions:ProjectTeamPermissions_Read  = 100
	// Permissions:ProjectTeamPermissions_Write = 200
	Permissions int64
}

type PipelineTeam struct {
	gorm.Model
	Created     int64 `json:"created"`
	Updated     int64 `json:"updated"`
	TeamId      int64
	PipelineId  int64
	// Permissions:PipelineTeamPermissions_Read  = 100
	// Permissions:PipelineTeamPermissions_Write = 200
	Permissions int64
}
