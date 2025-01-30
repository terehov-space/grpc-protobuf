package repositories

import (
	"fmt"
	"gorm.io/gorm"
	"wpc/user-service/internal/models"
)

type UserRepository struct {
	DB *gorm.DB
}

func (r *UserRepository) GetUserByID(id int64) (*models.User, error) {
	var user models.User

	result := r.DB.First(&user, "id = ?", id)

	if result.Error != nil {
		return nil, fmt.Errorf("user not found: %v", result.Error)
	}

	return &user, nil
}

func (r *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User

	result := r.DB.First(&user, "email = ?", email)

	if result.Error != nil {
		return nil, fmt.Errorf("user not found: %v", result.Error)
	}

	return &user, nil
}
