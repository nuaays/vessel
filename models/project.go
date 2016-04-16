package models

import "github.com/jinzhu/gorm"

type Project struct {
	gorm.Model
	WorkspaceID uint   `json:"workspaceID"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Actived     bool   `json:"actived"`
}

//
// func (pj *Project) Create(wid int64, name, description string) (int64, error) {
// 	return 0, nil
// }
//
// func (pj *Project) Put(id int64, name, description string) error {
// 	return nil
// }
//
// func (pj *Project) Get(id int64) (*Project, error) {
// 	return nil, nil
// }
//
// func (pj *Project) Delete(id int64) error {
// 	return nil
//
// }
