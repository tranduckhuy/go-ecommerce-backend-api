package entities

type Role struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	Name        string `gorm:"type:varchar(50);unique;not null"`
	Description string `gorm:"type:varchar(255)"`
}

func (r *Role) TableName() string {
	return "roles"
}
