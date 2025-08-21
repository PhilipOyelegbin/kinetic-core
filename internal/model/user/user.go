package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	IsVerified bool   `json:"is_verified" gorm:"default:false"`
	VerifyToken string `json:"verify_token" gorm:"default:null"`
	VerifyExpTime int64 `json:"verify_exp_time" gorm:"default:null"`
	ResetToken    string `json:"reset_token" gorm:"default:null"`
	ResetExpTime  int64  `json:"reset_exp_time" gorm:"default:null"`
}