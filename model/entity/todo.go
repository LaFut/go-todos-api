package entity

import (
	"github.com/jinzhu/gorm"
)

type InterfaceEntity interface {
}

type TodoFields struct {
	Name     string `bidding:"required" json:"name"`
	ParentId uint   `gorm:"INDEX" json:"parentid"`
}

type Todo struct {
	InterfaceEntity `json:"-"`
	gorm.Model
	*TodoFields
	Children []Todo
}
