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
	CreateDebt(in *pb.DebtsRequest) (*pb.Debts, error)
	GetDebt(in *pb.DebtsID) (*pb.Debts, error)
	GetListDebts(in *pb.FilterDebts) (*pb.DebtsList, error)
	GetClientDebts(in *pb.ClientID) (*pb.DebtsList, error)
	PayPayment(in *pb.PayDebtsReq) (*pb.Debts, error)
}

type PaymentsRepo interface {
	GetPayment(in *pb.PaymentID) (*pb.Payment, error)
	GetPaymentsByDebtId(in *pb.DebtsID) (*pb.PaymentList, error)
	GetPayments(in *pb.FilterPayment) (*pb.PaymentList, error)
}
