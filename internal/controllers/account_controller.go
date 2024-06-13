package controllers

import (
	"context"
	"log"

	"github.com/gurodrigues-dev/account-gRPC/internal/service"
	"github.com/gurodrigues-dev/account-gRPC/pb"
	"github.com/gurodrigues-dev/account-gRPC/types"
)

type Server struct {
	pb.UnimplementedAccountServer
	accountservice *service.AccountService
}

func NewAccountController(service *service.AccountService) *Server {
	return &Server{
		accountservice: service,
	}
}

func (s *Server) CreateBill(ctx context.Context, proto *pb.CreateBillRequest) (*pb.CreateBillResponse, error) {

	input := types.Bill{
		Name:        proto.GetName(),
		Description: proto.GetDescription(),
		Amount:      proto.GetAmount(),
		Installment: int64(proto.GetInstallment()),
	}

	err := input.IsInstallmentHealthy()
	if err != nil {
		log.Printf("installment is greater at 128: %s", err.Error())
		return nil, err
	}

	err = s.accountservice.CreateBill(ctx, &input)
	if err != nil {
		log.Printf("error to create bill: %s", err.Error())
		return nil, err
	}

	return &pb.CreateBillResponse{
		Name:        input.Name,
		Description: input.Description,
		Amount:      input.Amount,
		Installment: input.Installment,
	}, nil

}

func (s *Server) DeleteBill(ctx context.Context, proto *pb.DeleteBillRequest) (*pb.DeleteBillResponse, error) {

	input := types.Bill{
		ID: proto.GetId(),
	}

	err := s.accountservice.DeleteBill(ctx, &input.ID)
	if err != nil {
		log.Printf("error to delete bill: %s", err.Error())
		return nil, err
	}

	return &pb.DeleteBillResponse{
		Message: "bill deleted w/ successfully",
	}, nil
}
