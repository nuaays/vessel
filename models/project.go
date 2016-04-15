package models

import (
)

type Project struct {
	Id          int64     `json:"id"`
	WorkspaceId int64     `json:"workspaceId"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Actived     bool      `json:"actived"`
	Created     int64 `json:"created"`
	Updated     int64 `json:"updated"`
	Memo        string    `json:"memo"`
}

func (pj *Project) Create(wid int64, name, description string) (int64, error) {
	return 0, nil
}

func (pj *Project) Put(id int64, name, description string) error {
	return nil
}

func (pj *Project) Get(id int64) (*Project, error) {
	return nil, nil
}

func (pj *Project) Delete(id int64) error {
	return nil

}
