package usecase

import (
	"context"
	"debts-service/internal/generated/user"
	"debts-service/internal/usecase/webapi"
	"fmt"
	"github.com/shopspring/decimal"
	"log/slog"
	"strings"
	"time"

	pb "debts-service/internal/generated/debts"
)

type DebtsServiceServer struct {
	userService user.AuthServiceClient
	pb.UnimplementedDebtsServiceServer
	repo  DebtsRepo
	repo1 PaymentsRepo
	log   *slog.Logger
}

func NewDebtsServiceServer(repo DebtsRepo, repo1 PaymentsRepo, log *slog.Logger, userClient user.AuthServiceClient) *DebtsServiceServer {
	return &DebtsServiceServer{repo: repo, repo1: repo1, log: log, userService: userClient}
}

func (s *DebtsServiceServer) CreateDebts(ctx context.Context, in *pb.DebtsRequest) (*pb.Debts, error) {
	s.log.Info("CreateDebt called", "client_id", in.ClientId)

	debt, err := s.repo.CreateDebt(in)
	if err != nil {
		s.log.Error("Failed to create debt", "error", err)
		return nil, fmt.Errorf("could not create debt: %w", err)
	}

	return debt, nil
}

func (s *DebtsServiceServer) GetDebts(ctx context.Context, in *pb.DebtsID) (*pb.Debts, error) {
	s.log.Info("GetDebt called", "debt_id", in.Id)

	debt, err := s.repo.GetDebt(in)
	if err != nil {
		s.log.Error("Failed to retrieve debt", "debt_id", in.Id, "error", err)
		return nil, fmt.Errorf("could not retrieve debt: %w", err)
	}

	return debt, nil
}

func (s *DebtsServiceServer) PayDebts(ctx context.Context, in *pb.PayDebtsReq) (*pb.Debts, error) {
	s.log.Info("GetListDebts called", "filters", in)

	debts, err := s.repo.PayPayment(in)
	if err != nil {
		s.log.Error("Failed to retrieve debts list", "error", err)
		return nil, fmt.Errorf("could not retrieve debts list: %w", err)
	}

	return debts, nil
}

func (s *DebtsServiceServer) GetListDebts(ctx context.Context, in *pb.FilterDebts) (*pb.DebtsList, error) {
	s.log.Info("GetClientDebts called", "client_id", in.CompanyId)

	debts, err := s.repo.GetListDebts(in)
	if err != nil {
		s.log.Error("Failed to retrieve client debts", "client_id", in.CompanyId, "error", err)
		return nil, fmt.Errorf("could not retrieve client debts: %w", err)
	}

	return debts, nil
}

func (s *DebtsServiceServer) GetClientDebts(ctx context.Context, in *pb.ClientID) (*pb.DebtsList, error) {

	debt, err := s.repo.GetClientDebts(in)
	if err != nil {
		s.log.Error("Failed to pay debt", "debt_id", in.Id, "error", err)
		return nil, fmt.Errorf("could not pay debt: %w", err)
	}

	return debt, nil
}

func (s *DebtsServiceServer) GetPayment(ctx context.Context, in *pb.PaymentID) (*pb.Payment, error) {
	s.log.Info("GetPayment called", "payment_id", in.Id)

	payment, err := s.repo1.GetPayment(in)
	if err != nil {
		s.log.Error("Failed to retrieve payment", "payment_id", in.Id, "error", err)
		return nil, fmt.Errorf("could not retrieve payment: %w", err)
	}

	s.log.Info("Payment retrieved successfully", "payment_id", payment.Id)
	return payment, nil
}

func (s *DebtsServiceServer) GetPaymentsByDebtsId(ctx context.Context, in *pb.DebtsID) (*pb.PaymentList, error) {
	s.log.Info("GetPaymentsByDebtId called", "debt_id", in.Id)

	payments, err := s.repo1.GetPaymentsByDebtId(in)
	if err != nil {
		s.log.Error("Failed to retrieve payments by debt ID", "debt_id", in.Id, "error", err)
		return nil, fmt.Errorf("could not retrieve payments by debt ID: %w", err)
	}

	s.log.Info("Payments retrieved by debt ID successfully", "debt_id", in.Id, "count", len(payments.Payments))
	return payments, nil
}

func (s *DebtsServiceServer) GetPayments(ctx context.Context, in *pb.FilterPayment) (*pb.PaymentList, error) {
	s.log.Info("GetPayments called", "filters", in)

	payments, err := s.repo1.GetPayments(in)
	if err != nil {
		s.log.Error("Failed to retrieve payments list", "error", err)
		return nil, fmt.Errorf("could not retrieve payments list: %w", err)
	}

	s.log.Info("Payments list retrieved successfully", "count", len(payments.Payments))
	return payments, nil
}

func (s *DebtsServiceServer) GetTotalDebtSum(ctx context.Context, in *pb.CompanyID) (*pb.SumMoney, error) {
	s.log.Info("GetTotalDebtSum", "req", in)

	res, err := s.repo.GetTotalDebtSum(in)
	if err != nil {
		s.log.Error("Failed to retrieve total debt sum", "error", err)
		return nil, fmt.Errorf("could not retrieve total debt sum: %w", err)
	}

	return res, nil
}

func (s *DebtsServiceServer) GetUserTotalDebtSum(ctx context.Context, in *pb.ClientID) (*pb.SumMoney, error) {

	s.log.Info("GetUserTotalDebtSum", "req", in)

	res, err := s.repo.GetUserTotalDebtSum(in)
	if err != nil {
		s.log.Error("Failed to retrieve user total debt sum", "error", err)
		return nil, fmt.Errorf("could not retrieve user total debt sum: %w", err)
	}

	return res, nil
}

func (s *DebtsServiceServer) GetUserPayments(ctx context.Context, in *pb.ClientID) (*pb.UserPaymentsRes, error) {

	s.log.Info("GetUserPayments", "req", in)

	res, err := s.repo1.GetUserPayments(in)

	if err != nil {
		s.log.Error("Failed to retrieve user payments", "error", err)
		return nil, fmt.Errorf("could not retrieve user payments: %w", err)
	}

	return res, nil
}

func (s *DebtsServiceServer) GetDebtsForExel(ctx context.Context, in *pb.FilterExelDebt) (*pb.ExelDebtsList, error) {
	s.log.Info("GetDebtsForExel", "req", in)

	debts, err := s.repo.GetDebtsForExel(in)
	if err != nil {
		s.log.Error("Error getting debts for exel", "error", err)
		return nil, fmt.Errorf("failed to get debts for exel: %w", err)
	}

	usdRate, err := webapi.GetUSDCourse()
	if err != nil {
		s.log.Error("Error getting usd rate", "error", err)
		return nil, fmt.Errorf("failed to get USD rate: %w", err)
	}

	usdRateDec := decimal.NewFromFloat(usdRate)
	reqCurrency := strings.ToLower(in.Currency)

	var res pb.ExelDebtsList

	for _, i := range debts.Debts {
		clientData, err := s.userService.GetClient(context.Background(), &user.UserIDRequest{
			Id:        i.ClientID,
			CompanyId: in.CompanyId,
		})
		if err != nil {
			s.log.Error("Error getting user data", "error", err)
			return &res, fmt.Errorf("error in getting user data: %w", err)
		}

		totalAmount := decimal.Zero
		amountPaid := decimal.Zero
		debtBalance := decimal.Zero
		var lastPaidDate time.Time
		var lastPaidDateStr string

		for _, j := range i.Debts {

			if j.LastPaidDate != "" {
				parsedDate, err := time.Parse("2006-01-02", j.LastPaidDate)
				if err == nil && parsedDate.After(lastPaidDate) {
					lastPaidDate = parsedDate
					lastPaidDateStr = j.LastPaidDate
				}
			}

			jTotal := decimal.NewFromFloat(j.TotalAmount)
			jPaid := decimal.NewFromFloat(j.AmountPaid)
			jDebt := decimal.NewFromFloat(j.DebtBalance)

			jCurrency := strings.ToLower(j.CurrencyCode)

			if jCurrency == reqCurrency {
				totalAmount = totalAmount.Add(jTotal)
				amountPaid = amountPaid.Add(jPaid)
				debtBalance = debtBalance.Add(jDebt)
			} else if reqCurrency == "usd" {
				totalAmount = totalAmount.Add(jTotal.Div(usdRateDec))
				amountPaid = amountPaid.Add(jPaid.Div(usdRateDec))
				debtBalance = debtBalance.Add(jDebt.Div(usdRateDec))
			} else if reqCurrency == "uzs" {
				totalAmount = totalAmount.Add(jTotal.Mul(usdRateDec))
				amountPaid = amountPaid.Add(jPaid.Mul(usdRateDec))
				debtBalance = debtBalance.Add(jDebt.Mul(usdRateDec))
			}
		}

		var debt pb.ExelDebts
		debt.ClientId = i.ClientID
		debt.ClientFullName = clientData.FullName
		debt.ClientPhone = clientData.Phone

		debt.TotalAmount = totalAmount.Round(2).InexactFloat64()
		debt.AmountPaid = amountPaid.Round(2).InexactFloat64()
		debt.DebtsBalance = debtBalance.Round(2).InexactFloat64()

		debt.LastPaymentDate = lastPaidDateStr
		debt.CurrencyCode = in.Currency

		res.Debts = append(res.Debts, &debt)
	}

	return &res, nil
}
