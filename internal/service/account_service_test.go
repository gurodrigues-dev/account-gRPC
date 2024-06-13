package service

import (
	"context"
	"database/sql"
	"testing"

	"github.com/gurodrigues-dev/account-gRPC/config"
	"github.com/gurodrigues-dev/account-gRPC/internal/repository"
	"github.com/gurodrigues-dev/account-gRPC/types"
)

func newPostgres(dbConfig config.Database) string {
	return "user=" + dbConfig.User +
		" password=" + dbConfig.Password +
		" dbname=" + dbConfig.Name +
		" host=" + dbConfig.Host +
		" port=" + dbConfig.Port +
		" sslmode=disable"
}

func mockAccount() *types.Bill {
	return &types.Bill{
		Name:        "Logitech GPRO",
		Description: "Mouse Logitech GPRO Superlight X",
		Amount:      "80",
		Installment: 10,
	}
}

func setupTestDB(t *testing.T) (*sql.DB, *AccountService) {
	t.Helper()

	config, err := config.Load("../../config/config.yaml")
	if err != nil {
		t.Fatalf("falha ao carregar a configuração: %v", err)
	}

	db, err := sql.Open("postgres", newPostgres(config.Database))
	if err != nil {
		t.Fatalf("falha ao conectar ao banco de dados: %v", err)
	}

	accountRepository := repository.NewAccountRepository(db)
	accountService := NewAccountService(accountRepository)

	return db, accountService
}

func TestCreateBill(t *testing.T) {

	db, accountService := setupTestDB(t)
	defer db.Close()

	accountMock := mockAccount()

	err := accountService.CreateBill(context.Background(), accountMock)

	if err != nil {
		t.Errorf("Erro ao criar conta ou mensalidade: %v", err.Error())
	}

}

func TestDeleteBill(t *testing.T) {

	db, accountService := setupTestDB(t)
	defer db.Close()

	accountMock := mockAccount()

	err := accountService.DeleteBill(context.Background(), &accountMock.ID)

	if err != nil {
		t.Errorf("Erro ao deletar conta ou mensalidade: %v", err.Error())
	}
}
