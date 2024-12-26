package usecase

import (
	"context"
	"fmt"
	"log/slog"

	pb "debts-service/internal/generated/debts"
)

type ClientServiceServer struct {
	pb.UnimplementedClientServiceServer
	repo ClientsRepo
	log  *slog.Logger
}

func NewClientServiceServer(repo ClientsRepo, log *slog.Logger) *ClientServiceServer {
	return &ClientServiceServer{repo: repo, log: log}
}

func (s *ClientServiceServer) AddClient(ctx context.Context, req *pb.CreateClients) (*pb.Client, error) {
	s.log.Info("AddClient called", "full_name", req.FullName)

	client, err := s.repo.AddClient(req)
	if err != nil {
		s.log.Error("Failed to add client", "error", err)
		return nil, fmt.Errorf("could not add client: %w", err)
	}

	s.log.Info("Client added successfully", "client_id", client.Id)
	return client, nil
}

func (s *ClientServiceServer) GetClient(ctx context.Context, req *pb.ClientID) (*pb.Client, error) {
	s.log.Info("GetClient called", "client_id", req.Id)

	client, err := s.repo.GetClient(req)
	if err != nil {
		s.log.Error("Failed to retrieve client", "client_id", req.Id, "error", err)
		return nil, fmt.Errorf("could not retrieve client: %w", err)
	}

	s.log.Info("Client retrieved successfully", "client_id", client.Id)
	return client, nil
}

func (s *ClientServiceServer) UpdateClient(ctx context.Context, req *pb.ClientUpdate) (*pb.Client, error) {
	s.log.Info("UpdateClient called", "client_id", req.Id)

	client, err := s.repo.UpdateClient(req)
	if err != nil {
		s.log.Error("Failed to update client", "client_id", req.Id, "error", err)
		return nil, fmt.Errorf("could not update client: %w", err)
	}

	s.log.Info("Client updated successfully", "client_id", client.Id)
	return client, nil
}

func (s *ClientServiceServer) GetAllClients(ctx context.Context, req *pb.FilterClient) (*pb.ClientList, error) {
	s.log.Info("GetAllClients called", "filters", req)

	clients, err := s.repo.GetAllClients(req)
	if err != nil {
		s.log.Error("Failed to retrieve clients list", "error", err)
		return nil, fmt.Errorf("could not retrieve clients list: %w", err)
	}

	s.log.Info("Clients list retrieved successfully", "count", len(clients.Clients))
	return clients, nil
}
