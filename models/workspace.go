package models

import "github.com/jinzhu/gorm"

type Workspace struct {
	gorm.Model
	Name        string `json:"name"`
	Actived     bool   `json:"actived"`
	Description string `json:"description"`
	projects    []Project
}

// func (ws *Workspace) Demo() (int64, error) {
// 	return 0, nil
// }

//
// func (ws *Workspace) Create(name, description string) (int64, error) {
// 	return 0, nil
// }
//
// func (ws *Workspace) Put(id int64, name, description string) error {
// 	return nil
// }
//
// func (ws *Workspace) Get(id int64) (*Workspace, error) {
// 	return nil, nil
// }
//
// func (ws *Workspace) Delete(id int64) error {
// 	return nil
// }
