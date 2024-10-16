package models

// import "time"

type User struct {
	ID       string   `gorm:"primaryKey;unique" json:"id"`
	Email    string `gorm:"unique" json:"email"`
	Picture    string `gorm:"unique" json:"picture"`
	VerifiedEmail bool `gorm:"default:false" json:"verified_email"`
	Password []byte `json:"-"`
}
