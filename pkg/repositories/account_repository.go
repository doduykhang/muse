package repositories

import (
	"github.com/doduykhang/muse/pkg/models"
)

type AccountRepository interface {
	CrudRepository[models.Account, uint]
	FindAccountByEmail(email string) (*models.Account, error)
}

type accountRepository struct {
	CrudRepositoryImpl[models.Account, uint]
}

func (repository *accountRepository) FindAccountByEmail(email string) (*models.Account, error) {
	var account models.Account 
	result := db.First(&account, "email = ?", email)
	if result.Error != nil {
		return nil, result.Error
	}
	return &account, nil	
}

func NewAccountRepository() AccountRepository {
	return &accountRepository{}
}
