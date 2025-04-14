package repo

import (
	"database/sql"
	pb "debts-service/internal/generated/debts"
	"debts-service/internal/usecase"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
)

type paymentRepo struct {
	db *sqlx.DB
}

func NewPaymentRepo(db *sqlx.DB) usecase.PaymentsRepo {
	return &paymentRepo{db: db}
}

func (p *paymentRepo) GetPayment(in *pb.PaymentID) (*pb.Payment, error) {
	var payment pb.Payment

	query := `
		SELECT id, installment_id, payment_date, pay_type, payment_amount, created_at
		FROM payments
		WHERE id = $1
	`
	err := p.db.QueryRowx(query, in.Id).Scan(
		&payment.Id,
		&payment.InstallmentId,
		&payment.PaymentDate,
		&payment.PayType,
		&payment.PaymentAmount,
		&payment.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get payment: %w", err)
	}
	return &payment, nil
}

// GetPayments retrieves payments based on filter criteria
func (p *paymentRepo) GetPayments(in *pb.FilterPayment) (*pb.PaymentList, error) {
	query := `
		SELECT p.id, p.installment_id, p.payment_date, p.pay_type, p.payment_amount, p.created_at
		FROM payments p
		JOIN installment i ON p.installment_id = i.id
		WHERE i.company_id = $1
	`
	var args []interface{}
	args = append(args, in.CompanyId)

	if in.InstallmentId != "" {
		query += fmt.Sprintf(" AND p.installment_id = $%d", len(args)+1)
		args = append(args, in.InstallmentId)
	}
	if in.StartDate != "" {
		query += fmt.Sprintf(" AND p.payment_date >= $%d", len(args)+1)
		args = append(args, in.StartDate)
	}
	if in.EndDate != "" {
		query += fmt.Sprintf(" AND p.payment_date <= $%d", len(args)+1)
		args = append(args, in.EndDate)
	}

	rows, err := p.db.Queryx(query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	var payments []*pb.Payment
	for rows.Next() {
		var payment pb.Payment
		if err := rows.Scan(
			&payment.Id,
			&payment.InstallmentId,
			&payment.PaymentDate,
			&payment.PayType,
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
func (p *paymentRepo) GetPaymentsByDebtId(in *pb.PayDebtsID) (*pb.PaymentList, error) {
	query := `
        SELECT id, installment_id, payment_date, pay_type, payment_amount, created_at
        FROM payments
        WHERE installment_id = $1 AND pay_type = $2
        ORDER BY payment_date DESC`

	rows, err := p.db.Queryx(query, in.Id, in.PayType)
	if err != nil {
		return nil, fmt.Errorf("failed to get installment payments: %w", err)
	}
	defer rows.Close()

	var payments []*pb.Payment
	for rows.Next() {
		var payment pb.Payment
		if err := rows.Scan(
			&payment.Id,
			&payment.InstallmentId,
			&payment.PaymentDate,
			&payment.PayType,
			&payment.PaymentAmount,
			&payment.CreatedAt,
		); err != nil {
			return nil, fmt.Errorf("failed to scan payment: %w", err)
		}
		payments = append(payments, &payment)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over payments: %w", err)
	}

	return &pb.PaymentList{Payments: payments}, nil
}

func (p *paymentRepo) GetUserPayments(in *pb.ClientID) (*pb.UserPaymentsRes, error) {
	query := `
        SELECT 
            p.installment_id AS debt_id, 
            p.id AS payment_id, 
            p.payment_date, 
            p.payment_amount, 
            p.pay_type
        FROM payments p
        JOIN installment i ON i.id = p.installment_id 
        WHERE i.client_id = $1 AND i.company_id = $2 AND i.debt_type = $3
    `

	log.Println(in.Id, in.CompanyId, in.DebtType)

	rows, err := p.db.Queryx(query, in.Id, in.CompanyId, in.DebtType)
	if err != nil {
		return nil, fmt.Errorf("failed to query payments: %w", err)
	}
	defer rows.Close()

	var payments []*pb.Payments
	for rows.Next() {
		var payment pb.Payments
		if err := rows.Scan(
			&payment.DebtId,
			&payment.PaymentId,
			&payment.PaymentDate,
			&payment.PaymentAmount,
			&payment.PayType,
		); err != nil {
			return nil, fmt.Errorf("failed to scan payment: %w", err)
		}
		payments = append(payments, &payment)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %w", err)
	}

	return &pb.UserPaymentsRes{Payments: payments}, nil
}
