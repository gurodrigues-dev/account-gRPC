package service

import (
	"context"

	"github.com/gurodrigues-dev/account-gRPC/internal/repository"
	"github.com/gurodrigues-dev/account-gRPC/types"
)

type AccountService struct {
	accountrepository repository.IAccountRepository
}

func NewAccountService(repo repository.IAccountRepository) *AccountService {
	return &AccountService{
		accountrepository: repo,
	}
}

func (as *AccountService) CreateBill(ctx context.Context, bill *types.Bill) error {
	return as.accountrepository.CreateBill(ctx, bill)
}

func (as *AccountService) DeleteBill(ctx context.Context, id *string) error {
	return as.accountrepository.DeleteBill(ctx, id)
}
