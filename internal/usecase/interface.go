package usecase

import (
	pb "debts-service/internal/generated/debts"
	"debts-service/internal/usecase/entity"
)

type ClientsRepo interface {
	AddClient(in *pb.CreateClients) (*pb.Client, error)
	GetClient(in *pb.ClientID) (*pb.Client, error)
	UpdateClient(in *pb.ClientUpdate) (*pb.Client, error)
	GetAllClients(in *pb.FilterClient) (*pb.ClientList, error)
	CloseDebt(in *pb.ClientID) error
	OpenDebt(in *pb.ClientID) error
}

type DebtsRepo interface {
	CreateDebt(in *pb.DebtsRequest) (*pb.Debts, error)
	GetDebt(in *pb.DebtsID) (*pb.Debts, error)
	GetListDebts(in *pb.FilterDebts) (*pb.DebtsList, error)
	GetClientDebts(in *pb.ClientID) (*pb.DebtsList, error)
	PayPayment(in *pb.PayDebtsReq) (*pb.Debts, error)

	GetTotalDebtSum(in *pb.CompanyID) (*pb.SumMoney, error)
	GetUserTotalDebtSum(in *pb.ClientID) (*pb.SumMoney, error)

	GetDebtsForExel(in *pb.FilterExelDebt) (*entity.ListDebtsExelDb, error)
}

type PaymentsRepo interface {
	GetPayment(in *pb.PaymentID) (*pb.Payment, error)
	GetPaymentsByDebtId(in *pb.PayDebtsID) (*pb.PaymentList, error)
	GetPayments(in *pb.FilterPayment) (*pb.PaymentList, error)
	GetUserPayments(in *pb.ClientID) (*pb.UserPaymentsRes, error)
}
