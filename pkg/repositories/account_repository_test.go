package repositories_test

import (
	"testing"

	"github.com/doduykhang/muse/pkg/models"
	"github.com/doduykhang/muse/pkg/repositories"
)

var (
	testAccountRepository repositories.AccountRepository
)

func init() {
	testAccountRepository = repositories.NewAccountRepository()
}

func TestCreateAccount(t *testing.T) {
	account, err := testAccountRepository.Create(&models.Account{
		Email:    "test@gmail.com",
		Password: "testpassword",
		Active:   true,
		Banned:   true,
	})

	if err != nil {
		t.Errorf("CreateAccount() Create account return error: %v", err)
	}

	if account == nil && account.ID <= 0 {
		t.Errorf("CreateAccount() Created account id is: %d", account.ID)
	}

	t.Cleanup(func() {
		testAccountRepository.Delete(account.ID)
	})
}

func TestUpdateAccount(t *testing.T) {
	account, err := testAccountRepository.Create(&models.Account{
		Email:    "test@gmail.com",
		Password: "testpassword",
		Active:   true,
		Banned:   true,
	})

	if err != nil {
		t.Errorf("UpdateAccount() Create account return error: %v", err)
	}

	if account == nil && account.ID <= 0 {
		t.Errorf("UpdateAccount() Created account id is: %d", account.ID)
	}

	updatedAccount, err := testAccountRepository.Update(
		10,
		&models.Account{
			Email:    "test-update@gmail.com",
			Password: "testpassword",
			Active:   true,
			Banned:   true,
		},
	)

	if err != nil {
		t.Errorf("UpdateAccount() Update account return error: %v", err)
	}

	if updatedAccount == nil && updatedAccount.ID <= 0 {
		t.Errorf("UpdateAccount() Update account id is: %d", account.ID)
	}

	t.Cleanup(func() {
		testAccountRepository.Delete(account.ID)
	})
}

func TestDeleteAccount(t *testing.T) {
	account, err := testAccountRepository.Create(&models.Account{
		Email:    "test@gmail.com",
		Password: "testpassword",
		Active:   true,
		Banned:   true,
	})

	if err != nil {
		t.Errorf("DeleteAccount() Create account return error: %v", err)
	}

	if account == nil && account.ID <= 0 {
		t.Errorf("DeleteAccount() Created account id is: %d", account.ID)
	}

	err = testAccountRepository.Delete(account.ID)

	if err != nil {
		t.Errorf("DeleteAccount() Update account return error: %v", err)
	}

	t.Cleanup(func() {
		testAccountRepository.Delete(account.ID)
	})

}

func TestFindAccountById(t *testing.T) {
	account, err := testAccountRepository.Create(&models.Account{
		Email:    "test@gmail.com",
		Password: "testpassword",
		Active:   true,
		Banned:   true,
	})

	if err != nil {
		t.Errorf("FindAccountById() Create account return error: %v", err)
	}

	if account == nil && account.ID <= 0 {
		t.Errorf("FindAccountById() Created account id is: %d", account.ID)
	}

	_, err = testAccountRepository.FindById(account.ID)

	if err != nil {
		t.Errorf("FindAccountById() Update account return error: %v", err)
	}

	t.Cleanup(func() {
		testAccountRepository.Delete(account.ID)
	})
}

func TestFindAllAccount(t *testing.T) {
	account1, err := testAccountRepository.Create(&models.Account{
		Email:    "test@gmail.com",
		Password: "testpassword",
		Active:   true,
		Banned:   true,
	})

	if err != nil {
		t.Errorf("FindAccountById() Create account return error: %v", err)
	}

	account2, err := testAccountRepository.Create(&models.Account{
		Email:    "test2@gmail.com",
		Password: "testpassword",
		Active:   true,
		Banned:   true,
	})

	if err != nil {
		t.Errorf("FindAccountById() Create account return error: %v", err)
	}

	account3, err := testAccountRepository.Create(&models.Account{
		Email:    "test3@gmail.com",
		Password: "testpassword",
		Active:   true,
		Banned:   true,
	})

	if err != nil {
		t.Errorf("FindAccountById() Create account return error: %v", err)
	}

	accounts, err := testAccountRepository.FindAll()

	if err != nil {
		t.Errorf("FindAllAccount() return error: %v", err)
	}

	if len(accounts) != 3 {
		t.Errorf("FindAllAccount() return %d account(s), expect 3", len(accounts))
	}

	t.Cleanup(func() {
		testAccountRepository.Delete(account1.ID)
		testAccountRepository.Delete(account2.ID)
		testAccountRepository.Delete(account3.ID)
	})
}
