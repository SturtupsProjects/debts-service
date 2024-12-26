package app

import (
	"debts-service/config"
	pb "debts-service/internal/generated/debts"
	"debts-service/internal/usecase"
	"debts-service/internal/usecase/repo"
	"debts-service/pkg/logger"
	"debts-service/pkg/postgres"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

func Run(cfg *config.Config) {
	logs := logger.NewLogger()

	db, err := postgres.Connection(*cfg)
	if err != nil {
		log.Fatal(err)
	}

	clientRepo := repo.NewClientRepo(db)
	debtRepo := repo.NewInstallmentRepo(db)
	paymentRepo := repo.NewPaymentRepo(db)

	clientService := usecase.NewClientServiceServer(clientRepo, logs)
	debtService := usecase.NewDebtsServiceServer(debtRepo, logs)
	paymentService := usecase.NewPaymentServiceServer(paymentRepo, logs)

	listen, err := net.Listen("tcp", cfg.RUN_PORT)
	fmt.Println("Listening on port " + cfg.RUN_PORT)
	if err != nil {
		logs.Error("Error listening on port " + cfg.RUN_PORT)
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterDebtsServiceServer(grpcServer, debtService)
	pb.RegisterClientServiceServer(grpcServer, clientService)
	pb.RegisterPaymentServiceServer(grpcServer, paymentService)
	if err := grpcServer.Serve(listen); err != nil {
		logs.Error("Error starting server")
		log.Fatal(err)
	}
}
