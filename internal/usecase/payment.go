package usecase

import (
	"context"
	"fmt"
	"log/slog"

	pb "debts-service/internal/generated/debts"
)

type PaymentServiceServer struct {
	pb.UnimplementedPaymentServiceServer
	repo PaymentsRepo
	log  *slog.Logger
}

func NewPaymentServiceServer(repo PaymentsRepo, log *slog.Logger) *PaymentServiceServer {
	return &PaymentServiceServer{repo: repo, log: log}
}

func (s *PaymentServiceServer) GetPayment(ctx context.Context, req *pb.PaymentID) (*pb.Payment, error) {
	s.log.Info("GetPayment called", "payment_id", req.Id)

	payment, err := s.repo.GetPayment(req)
	if err != nil {
		s.log.Error("Failed to retrieve payment", "payment_id", req.Id, "error", err)
		return nil, fmt.Errorf("could not retrieve payment: %w", err)
	}

	s.log.Info("Payment retrieved successfully", "payment_id", payment.Id)
	return payment, nil
}

func (s *PaymentServiceServer) GetPaymentsByDebtId(ctx context.Context, req *pb.DebtID) (*pb.PaymentList, error) {
	s.log.Info("GetPaymentsByDebtId called", "debt_id", req.Id)

	payments, err := s.repo.GetPaymentsByDebtId(req)
	if err != nil {
		s.log.Error("Failed to retrieve payments by debt ID", "debt_id", req.Id, "error", err)
		return nil, fmt.Errorf("could not retrieve payments by debt ID: %w", err)
	}

	s.log.Info("Payments retrieved by debt ID successfully", "debt_id", req.Id, "count", len(payments.Payments))
	return payments, nil
}

func (s *PaymentServiceServer) GetPayments(ctx context.Context, req *pb.FilterPayment) (*pb.PaymentList, error) {
	s.log.Info("GetPayments called", "filters", req)

	payments, err := s.repo.GetPayments(req)
	if err != nil {
		s.log.Error("Failed to retrieve payments list", "error", err)
		return nil, fmt.Errorf("could not retrieve payments list: %w", err)
	}

	s.log.Info("Payments list retrieved successfully", "count", len(payments.Payments))
	return payments, nil
}
