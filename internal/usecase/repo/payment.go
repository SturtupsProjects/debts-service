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

// PayPayment records a payment for an installment
func (p *paymentRepo) PayPayment(in *pb.PayDebtReq) (*pb.Debt, error) {
	tx, err := p.db.Beginx()
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	// Insert payment record
	paymentQuery := `
        INSERT INTO payments (id, installment_id, payment_amount, payment_date)
        VALUES (gen_random_uuid(), $1, $2, CURRENT_TIMESTAMP)`

	_, err = tx.Exec(paymentQuery, in.DebtId, in.PaidAmount)
	if err != nil {
		return nil, fmt.Errorf("failed to insert payment: %w", err)
	}

	// Update installment record
	var debt pb.Debt
	debtQuery := `
        UPDATE installment
        SET amount_paid = amount_paid + $1,
            last_payment_date = CURRENT_TIMESTAMP,
            present_month = present_month + 1,
            is_fully_paid = CASE 
                WHEN amount_paid + $1 >= total_amount THEN true 
                ELSE false 
            END
        WHERE id = $2
        RETURNING id, months_duration, present_month, total_amount, amount_paid, last_payment_date, is_fully_paid, created_at`

	err = tx.QueryRowx(debtQuery, in.PaidAmount, in.DebtId).Scan(
		&debt.Id,
		&debt.MonthsDuration,
		&debt.PresentMonth,
		&debt.TotalAmount,
		&debt.AmountPaid,
		&debt.LastPaymentDate,
		&debt.IsFullyPaid,
		&debt.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("debt not found: %w", err)
		}
		return nil, fmt.Errorf("failed to update installment: %w", err)
	}

	if err = tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return &debt, nil
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
