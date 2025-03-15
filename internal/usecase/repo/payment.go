package repo

import (
	"database/sql"
	pb "debts-service/internal/generated/debts"
	"debts-service/internal/usecase"
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
)

type paymentRepo struct {
	db *sqlx.DB
}

func NewPaymentRepo(db *sqlx.DB) usecase.PaymentsRepo {
	return &paymentRepo{db: db}
}

// GetPayment retrieves a specific payment by its ID
func (p *paymentRepo) GetPayment(in *pb.PaymentID) (*pb.Payment, error) {
	var payment pb.Payment

	query := `
        SELECT id, installment_id, payment_date, payment_amount, created_at
        FROM payments
        WHERE id = $1`

	// Use QueryRowx for a single result
	err := p.db.QueryRowx(query, in.Id).Scan(
		&payment.Id,
		&payment.InstallmentId,
		&payment.PaymentDate,
		&payment.PaymentAmount,
		&payment.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("payment not found: %w", err)
		}
		return nil, fmt.Errorf("failed to get payment: %w", err)
	}

	return &payment, nil
}

// GetPayments retrieves payments based on filter criteria
func (p *paymentRepo) GetPayments(in *pb.FilterPayment) (*pb.PaymentList, error) {
	var payments []*pb.Payment
	baseQuery := `
		SELECT id, installment_id, payment_date, payment_amount, created_at
		FROM payments
		WHERE company_id = $1` // Жёсткое условие для company_id
	var args []interface{}
	args = append(args, in.CompanyId) // Первым параметром всегда будет company_id

	var conditions []string

	if in.InstallmentId != "" {
		conditions = append(conditions, fmt.Sprintf("installment_id = $%d", len(args)+1))
		args = append(args, in.InstallmentId)
	}

	if in.StartDate != "" {
		conditions = append(conditions, fmt.Sprintf("payment_date >= $%d", len(args)+1))
		args = append(args, in.StartDate)
	}

	if in.EndDate != "" {
		conditions = append(conditions, fmt.Sprintf("payment_date <= $%d", len(args)+1))
		args = append(args, in.EndDate)
	}

	// Если есть дополнительные условия, добавляем их в запрос
	if len(conditions) > 0 {
		baseQuery += " AND " + strings.Join(conditions, " AND ")
	}

	// Выполнение запроса
	rows, err := p.db.Queryx(baseQuery, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	// Чтение строк из результата
	for rows.Next() {
		var payment pb.Payment
		if err := rows.Scan(
			&payment.Id,
			&payment.InstallmentId,
			&payment.PaymentDate,
			&payment.PaymentAmount,
			&payment.CreatedAt,
		); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		payments = append(payments, &payment)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error while iterating rows: %w", err)
	}

	return &pb.PaymentList{Payments: payments}, nil
}

// GetPaymentsByDebtId retrieves all payments for a specific debt installment
func (p *paymentRepo) GetPaymentsByDebtId(in *pb.DebtsID) (*pb.PaymentList, error) {
	var payments []*pb.Payment

	query := `
        SELECT id, installment_id, payment_date, payment_amount, created_at
        FROM payments
        WHERE installment_id = $1
        ORDER BY payment_date DESC`

	// Execute the query with the provided installment ID
	rows, err := p.db.Queryx(query, in.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to get installment payments: %w", err)
	}
	defer rows.Close()

	// Iterate over the result set
	for rows.Next() {
		var payment pb.Payment
		if err := rows.Scan(
			&payment.Id,
			&payment.InstallmentId,
			&payment.PaymentDate,
			&payment.PaymentAmount,
			&payment.CreatedAt,
		); err != nil {
			return nil, fmt.Errorf("failed to scan payment: %w", err)
		}
		payments = append(payments, &payment)
	}

	// Check for errors after iterating over rows
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over payments: %w", err)
	}

	return &pb.PaymentList{Payments: payments}, nil
}

func (p *paymentRepo) GetUserPayments(in *pb.ClientID) (*pb.UserPaymentsRes, error) {
	const query = `
        SELECT p.installment_id, p.id, p.payment_date, p.payment_amount 
        FROM payments p
        JOIN installment i ON i.id = p.installment_id 
        WHERE i.client_id = $1 AND i.company_id = $2
    `

	rows, err := p.db.Query(query, in.Id, in.CompanyId)
	if err != nil {
		return nil, fmt.Errorf("failed to query payments: %w", err)
	}
	defer rows.Close()

	var payments []*pb.Payments

	for rows.Next() {
		var payment pb.Payments
		if err := rows.Scan(&payment.DebtId, &payment.PaymentId, &payment.PaymentDate, &payment.PaymentAmount); err != nil {
			return nil, fmt.Errorf("failed to scan payment: %w", err)
		}
		payments = append(payments, &payment)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %w", err)
	}

	return &pb.UserPaymentsRes{Payments: payments}, nil
}
