package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(255);column:Username" json:"username"`
	Password string `gorm:"type:varchar(255);column:Password" json:"-"`
	Role     string `gorm:"type:varchar(255);column:Role" json:"role"`
	Todos    []Todo `gorm:"foreignkey:UserID" json:"todos"`
}

func (User) TableName() string {
	return "users"
}
