package models
import "time"

type User struct {
    ID          uint      `gorm:"primaryKey"`
    Email       string    `gorm:"uniqueIndex;size:255"`
    Password    string    `gorm:"size:255"`
    Name        string    `gorm:"size:255"`
    Phone       string    `gorm:"size:255"`         
    Address     string    `gorm:"size:255"`
    StoreStatus bool      `gorm:"default:false"`   
    CreatedAt   time.Time
    UpdatedAt   time.Time
     Store       *Store    `json:"store"`
}


type LoginRequest struct {
    Email    string `json:"email" binding:"required"`
    Password string `json:"password" binding:"required"`
}

type RegisterRequest struct {
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required,min=8"` 
    Name     string `json:"name" binding:"required"` 
    Phone    string `json:"phone" binding:"required"` 
    Address   string `json:"address" binding:"required"`
}

type UserProfile struct {
    ID    int    `json:"id"`
    Email string `json:"email"`
    Name  string `json:"name"`
}

type EditProfileRequest struct{
    Name    string `json:"name"`
    Phone   string `json:"phone"`
    Address string `json:"address"`
    Email string `json:"email"`
}
