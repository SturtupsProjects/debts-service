package usecase

import pb "debts-service/internal/generated/debts"

type ClientsRepo interface {
	AddClient(in *pb.CreateClients) (*pb.Client, error)
	GetClient(in *pb.ClientID) (*pb.Client, error)
	UpdateClient(in *pb.ClientUpdate) (*pb.Client, error)
	GetAllClients(in *pb.FilterClient) (*pb.ClientList, error)
	CloseDebt(in *pb.ClientID) error
	OpenDebt(in *pb.ClientID) error
}

type DebtsRepo interface {
	CreateDebt(in *pb.DebtRequest) (*pb.Debt, error)
	GetDebt(in *pb.DebtID) (*pb.Debt, error)
	GetListDebts(in *pb.FilterDebt) (*pb.DebtsList, error)
	GetClientDebts(in *pb.ClientID) (*pb.DebtsList, error)
}

type PaymentsRepo interface {
	PayPayment(in *pb.PayDebtReq) (*pb.Debt, error)
	GetPayment(in *pb.PaymentID) (*pb.Payment, error)
	GetPaymentsByDebtId(in *pb.DebtID) (*pb.PaymentList, error)
	GetPayments(in *pb.FilterPayment) (*pb.PaymentList, error)
}
