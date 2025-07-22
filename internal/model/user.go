package model

type User struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Name  string `json:"name"`
	Email string `gorm:"unique" json:"email"`
}

func (User) TableName() string {
	return "users"
}
