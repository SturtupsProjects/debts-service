package app

import (
	"debts-service/config"
	pb "debts-service/internal/generated/debts"
	"debts-service/internal/generated/user"
	"debts-service/internal/usecase"
	"debts-service/internal/usecase/repo"
	"debts-service/pkg/logger"
	"debts-service/pkg/postgres"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
)

func Run(cfg *config.Config) {
	logs := logger.NewLogger()

	db, err := postgres.Connection(*cfg)
	if err != nil {
		log.Fatal(err)
	}

	debtRepo := repo.NewInstallmentRepo(db)
	paymentRepo := repo.NewPaymentRepo(db)

	connUserService, err := grpc.NewClient("crm-admin_auth"+cfg.USER_SERVICE_PORT, grpc.WithTransportCredentials(insecure.NewCredentials()))
	userServiceClient := user.NewAuthServiceClient(connUserService)

	debtService := usecase.NewDebtsServiceServer(debtRepo, paymentRepo, logs, userServiceClient)

	listen, err := net.Listen("tcp", cfg.RUN_PORT)
	fmt.Println("Listening on port " + cfg.RUN_PORT)
	if err != nil {
		logs.Error("Error listening on port " + cfg.RUN_PORT)
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterDebtsServiceServer(grpcServer, debtService)
	if err := grpcServer.Serve(listen); err != nil {
		logs.Error("Error starting server")
		log.Fatal(err)
	}
}
