package models

type Acc struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Name     string `json:"name"`
	Position string `json:"position"`
}
