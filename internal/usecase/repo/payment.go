package repo

import (
	"database/sql"
	pb "debts-service/internal/generated/debts"
	"debts-service/internal/usecase"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type paymentRepo struct {
	db *sqlx.DB
}

func NewPaymentRepo(db *sqlx.DB) usecase.PaymentsRepo {
	return &paymentRepo{db: db}
}


// GetPayment retrieves a specific payment
func (p *paymentRepo) GetPayment(in *pb.PaymentID) (*pb.Payment, error) {
	var payment pb.Payment

	query := `
        SELECT id, installment_id, payment_date, payment_amount, created_at
        FROM payments
        WHERE id = $1`

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
// GetPayments retrieves payments based on filter criteria
func (p *paymentRepo) GetPayments(in *pb.FilterPayment) (*pb.PaymentList, error) {
	var payments []*pb.Payment

	query := `
        SELECT id, installment_id, payment_date, payment_amount, created_at
        FROM payments
        WHERE (COALESCE($1, '') = '' OR installment_id = $1)
          AND (COALESCE($2, '') = '' OR payment_date >= $2::timestamp)
          AND (COALESCE($3, '') = '' OR payment_date <= $3::timestamp)
        ORDER BY payment_date DESC`

	rows, err := p.db.Queryx(query, in.InstallmentId, in.StartDate, in.EndDate)
	if err != nil {
		return nil, fmt.Errorf("failed to get payments: %w", err)
	}
	defer rows.Close()

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

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over payments: %w", err)
	}

	return &pb.PaymentList{Payments: payments}, nil
}

// GetPaymentsByDebtId retrieves all payments for a specific installment
func (p *paymentRepo) GetPaymentsByDebtId(in *pb.DebtID) (*pb.PaymentList, error) {
	var payments []*pb.Payment

	query := `
        SELECT id, installment_id, payment_date, payment_amount, created_at
        FROM payments
        WHERE installment_id = $1
        ORDER BY payment_date DESC`

	rows, err := p.db.Queryx(query, in.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to get installment payments: %w", err)
	}
	defer rows.Close()

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

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over payments: %w", err)
	}

	return &pb.PaymentList{Payments: payments}, nil
}
