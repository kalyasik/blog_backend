package models

type Posts struct {
	ID     int64  `gorm:"primary_key;auto_increment;not_null" json:"id"`
	Title  string `gorm:"size:30;" json:"title" validate:"required,min=2,max=30"`
	Body   string `gorm:"not null" json:"body" validate:"required,min=5"`
	UserID int64  `gorm:"not null" json:"user_id"`
}
