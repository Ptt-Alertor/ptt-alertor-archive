package user

import (
	"time"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	gorm.Model
	Enable bool `json:"enable" gorm:"not null;index"`
	Profile {
		Account string `json:"account" gorm:"unique_index"`
		Email string `json:"email"`
		Line string `json:"line"`
		LineAccessToken string `json:"lineAccessToken"`
		Messenger string `json:"messenger"`
		Telegram string `json:"telegram"`
		TelegramChat int `json:"telegramChat"`
	} `json:"Profile"`
	Subscriptions []Subscription `json:"subscribes"`
	CreatedAt time.Time `json:"createTime"`
	UpdatedAt time.Time `json:"updateTime"`
}

type Subscription struct{
	Board string `json:"board"`
	Keywords []string `json:"keywords"`
	Authors []string `json:"authors"`
	Articles []string `json:"articles"`
	PushSum struct{
		Up uint `json:"up"`
		Down uint `json:"down"`
	} `json:"pushSum"`
}