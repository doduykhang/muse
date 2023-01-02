package repositories

import (
	"github.com/doduykhang/muse/pkg/models"
)

type UserRepository interface {
	CrudRepository[models.User, uint]
	FindUserByAccountID(id uint) (*models.User, error)
}

type userRepository struct {
	CrudRepositoryImpl[models.User, uint]
}

func (repository *userRepository) FindUserByAccountID(accountId uint) (*models.User, error) {
	var user models.User
	result := db.First(&user, "account_id = ?", accountId)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func NewUserRepository() UserRepository {
	return &userRepository{}
}
