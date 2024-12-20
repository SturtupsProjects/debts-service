package repo

import (
	"database/sql"
	"fmt"
	"strings"

	pb "debts-service/internal/generated/debts"
	"debts-service/internal/usecase"

	"github.com/jmoiron/sqlx"
)

type clientRepo struct {
	db *sqlx.DB
}

func NewClientRepo(db *sqlx.DB) usecase.ClientsRepo {
	return &clientRepo{db: db}
}

// AddClient inserts a new client into the database
func (c *clientRepo) AddClient(in *pb.CreateClients) (*pb.Client, error) {
	var client pb.Client

	query := `
		INSERT INTO clients (id, full_name, phone_number, address, telegram_username, notes)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, full_name, phone_number, address, telegram_username, telegram_user_id, has_debt, client_status, notes, created_at`

	err := c.db.QueryRowx(query, in.Id, in.FullName, in.PhoneNumber, in.Address, in.TelegramUsername, in.Notes).Scan(
		&client.Id,
		&client.FullName,
		&client.PhoneNumber,
		&client.Address,
		&client.TelegramUsername,
		&client.TelegramUserId,
		&client.HasDebt,
		&client.ClientStatus,
		&client.Notes,
		&client.CreatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to add client: %w", err)
	}

	return &client, nil
}

// GetClient retrieves a client by ID
func (c *clientRepo) GetClient(in *pb.ClientID) (*pb.Client, error) {
	var client pb.Client

	query := `
		SELECT id, full_name, phone_number, address, telegram_username, telegram_user_id, has_debt, client_status, notes, created_at
		FROM clients
		WHERE id = $1`

	err := c.db.QueryRowx(query, in.Id).Scan(
		&client.Id,
		&client.FullName,
		&client.PhoneNumber,
		&client.Address,
		&client.TelegramUsername,
		&client.TelegramUserId,
		&client.HasDebt,
		&client.ClientStatus,
		&client.Notes,
		&client.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("client not found")
		}
		return nil, fmt.Errorf("failed to get client: %w", err)
	}

	return &client, nil
}

// UpdateClient updates client details
func (c *clientRepo) UpdateClient(in *pb.ClientUpdate) (*pb.Client, error) {
	var client pb.Client

	query := `
		UPDATE clients
		SET full_name = $1, phone_number = $2, address = $3, has_debt = $4, notes = $5
		WHERE id = $6
		RETURNING id, full_name, phone_number, address, telegram_username, telegram_user_id, has_debt, client_status, notes, created_at`

	err := c.db.QueryRowx(query, in.FullName, in.PhoneNumber, in.Address, in.HasDebt, in.Notes, in.Id).Scan(
		&client.Id,
		&client.FullName,
		&client.PhoneNumber,
		&client.Address,
		&client.TelegramUsername,
		&client.TelegramUserId,
		&client.HasDebt,
		&client.ClientStatus,
		&client.Notes,
		&client.CreatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to update client: %w", err)
	}

	return &client, nil
}

// GetAllClients retrieves a list of clients with optional filtering
func (c *clientRepo) GetAllClients(in *pb.FilterClient) (*pb.ClientList, error) {
	var clients []*pb.Client
	var args []interface{}
	var filters []string
	argIndex := 1

	query := `SELECT id, full_name, phone_number, address, telegram_username, telegram_user_id, has_debt, client_status, notes, created_at FROM clients`

	if in.FullName != "" {
		filters = append(filters, fmt.Sprintf("full_name ILIKE $%d", argIndex))
		args = append(args, "%"+in.FullName+"%")
		argIndex++
	}
	if in.PhoneNumber != "" {
		filters = append(filters, fmt.Sprintf("phone_number = $%d", argIndex))
		args = append(args, in.PhoneNumber)
		argIndex++
	}
	if len(filters) > 0 {
		query += " WHERE " + strings.Join(filters, " AND ")
	}

	rows, err := c.db.Queryx(query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to query clients: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var client pb.Client
		if err := rows.Scan(
			&client.Id,
			&client.FullName,
			&client.PhoneNumber,
			&client.Address,
			&client.TelegramUsername,
			&client.TelegramUserId,
			&client.HasDebt,
			&client.ClientStatus,
			&client.Notes,
			&client.CreatedAt,
		); err != nil {
			return nil, fmt.Errorf("failed to scan client: %w", err)
		}
		clients = append(clients, &client)
	}

	return &pb.ClientList{Clients: clients}, nil
}

// CloseDebt marks a client's debt as closed
func (c *clientRepo) CloseDebt(in *pb.ClientID) error {
	query := `
		UPDATE clients
		SET has_debt = false
		WHERE id = $1`

	_, err := c.db.Exec(query, in.Id)
	if err != nil {
		return fmt.Errorf("failed to close debt: %w", err)
	}

	return nil
}

func (c *clientRepo) OpenDebt(in *pb.ClientID) error {
	query := `
		UPDATE clients
		SET has_debt = false
		WHERE id = $1`

	_, err := c.db.Exec(query, in.Id)
	if err != nil {
		return fmt.Errorf("failed to close debt: %w", err)
	}

	return nil
}
