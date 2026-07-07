package models

type Menu struct {
	ID          uint
	Name        string
	Description string
	Foods       []Food `gorm:"many2many:menu_foods;"`
}
