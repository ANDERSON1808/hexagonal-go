package db

import (
	"ANDERSON1808/hexagonal-go/internal/domain"
	"time"

	"gorm.io/gorm"
)

type UserEntity struct {
	ID        uint           `gorm:"primaryKey"`
	Name      string         `gorm:"type:varchar(100);not null"`
	Email     string         `gorm:"unique;not null"`
	Active    bool           `gorm:"default:true"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (u *UserEntity) ToDomain() *domain.User {
	return &domain.User{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		Active:    u.Active,
		CreatedAt: u.CreatedAt,
	}
}

func FromDomain(user *domain.User) *UserEntity {
	return &UserEntity{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Active:    user.Active,
		CreatedAt: user.CreatedAt,
	}
}
