package models

type Users struct {
	ID       int64  `gorm:"primary_key;auto_increment;not_null" json:"id"`
	Name     string `gorm:"size:80" json:"name"`
	Email    string `gorm:"size:60" json:"email"`
	Username string `gorm:"size:80;unique" json:"username"`
	Password string `gorm:"size:80" json:"password"`
}
