package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	*gorm.Model
	UserId   uuid.UUID `gorm:"type:char(36);primaryKey;index:idx_user_id"`
	Username string    `gorm:"type:varchar(50);unique;not null"`
	IsActive bool      `gorm:"type:bool;default:true"`
	Roles    []Role    `gorm:"many2many:go_user_roles;"`
}

func (u *User) TableName() string {
	return "users"
}
