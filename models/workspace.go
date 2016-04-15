package models

type Workspace struct {
	Id          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Actived     bool      `json:"actived"`
	Created     int64 `json:"created"`
	Updated     int64 `json:"updated"`
	Memo        string    `json:"memo"`
}

func (ws *Workspace) Create(name, description string) (int64, error) {
	return 0, nil
}

func (ws *Workspace) Put(id int64, name, description string) error {
	return nil
}

func (ws *Workspace) Get(id int64) (*Workspace, error) {
	return nil, nil
}

func (ws *Workspace) Delete(id int64) error {
	return nil
}
