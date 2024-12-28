package usecase

import (
	"context"
	"fmt"
	"log/slog"

	pb "debts-service/internal/generated/debts"
)

type DebtsServiceServer struct {
	pb.UnimplementedDebtsServiceServer
	repo DebtsRepo
	log  *slog.Logger
}

func NewDebtsServiceServer(repo DebtsRepo, log *slog.Logger) *DebtsServiceServer {
	return &DebtsServiceServer{repo: repo, log: log}
}

func (s *DebtsServiceServer) CreateDebt(ctx context.Context, req *pb.DebtRequest) (*pb.Debt, error) {
	s.log.Info("CreateDebt called", "client_id", req.ClientId)

	debt, err := s.repo.CreateDebt(req)
	if err != nil {
		s.log.Error("Failed to create debt", "error", err)
		return nil, fmt.Errorf("could not create debt: %w", err)
	}

	s.log.Info("Debt created successfully", "debt_id", debt.Id)
	return debt, nil
}

func (s *DebtsServiceServer) GetDebt(ctx context.Context, req *pb.DebtID) (*pb.Debt, error) {
	s.log.Info("GetDebt called", "debt_id", req.Id)

	debt, err := s.repo.GetDebt(req)
	if err != nil {
		s.log.Error("Failed to retrieve debt", "debt_id", req.Id, "error", err)
		return nil, fmt.Errorf("could not retrieve debt: %w", err)
	}

	s.log.Info("Debt retrieved successfully", "debt_id", debt.Id)
	return debt, nil
}

func (s *DebtsServiceServer) GetListDebts(ctx context.Context, req *pb.FilterDebt) (*pb.DebtsList, error) {
	s.log.Info("GetListDebts called", "filters", req)

	debts, err := s.repo.GetListDebts(req)
	if err != nil {
		s.log.Error("Failed to retrieve debts list", "error", err)
		return nil, fmt.Errorf("could not retrieve debts list: %w", err)
	}

	s.log.Info("Debts list retrieved successfully", "count", len(debts.Debts))
	return debts, nil
}

func (s *DebtsServiceServer) GetClientDebts(ctx context.Context, req *pb.ClientID) (*pb.DebtsList, error) {
	s.log.Info("GetClientDebts called", "client_id", req.Id)

	debts, err := s.repo.GetClientDebts(req)
	if err != nil {
		s.log.Error("Failed to retrieve client debts", "client_id", req.Id, "error", err)
		return nil, fmt.Errorf("could not retrieve client debts: %w", err)
	}

	s.log.Info("Client debts retrieved successfully", "client_id", req.Id, "count", len(debts.Debts))
	return debts, nil
}
func (s *DebtsServiceServer) PayDebt(ctx context.Context, req *pb.PayDebtReq) (*pb.Debt, error) {
	s.log.Info("PayDebt called", "debt_id", req.DebtId)

	debt, err := s.repo.PayPayment(req)
	if err != nil {
		s.log.Error("Failed to pay debt", "debt_id", req.DebtId, "error", err)
		return nil, fmt.Errorf("could not pay debt: %w", err)
	}

	s.log.Info("Debt paid successfully", "debt_id", debt.Id)
	return debt, nil
}
