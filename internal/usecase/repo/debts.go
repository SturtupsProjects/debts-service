package repo

import (
	"database/sql"
	"fmt"
	"strings"

	pb "debts-service/internal/generated/debts"
	"debts-service/internal/usecase"

	"github.com/jmoiron/sqlx"
)

type installmentRepo struct {
	db *sqlx.DB
}

func NewInstallmentRepo(db *sqlx.DB) usecase.DebtsRepo {
	return &installmentRepo{db: db}
}

// CreateDebt creates a new installment record
func (d *installmentRepo) CreateDebt(in *pb.DebtRequest) (*pb.Debt, error) {
	var debt pb.Debt

	query := `
        INSERT INTO installment (id, months_duration, client_id, total_amount, present_month, amount_paid)
        VALUES (gen_random_uuid(), $1, $2, $3, 1, 0)
        RETURNING id, months_duration, present_month, total_amount, amount_paid, last_payment_date, is_fully_paid, created_at`

	err := d.db.QueryRowx(query, in.MonthsDuration, in.ClientId, in.TotalAmount).Scan(
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
		return nil, fmt.Errorf("failed to create installment: %w", err)
	}

	return &debt, nil
}

// GetDebt retrieves an installment by ID
func (d *installmentRepo) GetDebt(in *pb.DebtID) (*pb.Debt, error) {
	var debt pb.Debt

	query := `
        SELECT id, months_duration, present_month, total_amount, amount_paid, last_payment_date, is_fully_paid, created_at
        FROM installment
        WHERE id = $1`

	err := d.db.QueryRowx(query, in.Id).Scan(
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
			return nil, fmt.Errorf("installment not found")
		}
		return nil, fmt.Errorf("failed to get installment: %w", err)
	}

	return &debt, nil
}

// GetListDebts retrieves a filtered list of installments
func (d *installmentRepo) GetListDebts(in *pb.FilterDebt) (*pb.DebtsList, error) {
	var debts []*pb.Debt
	var args []interface{}
	var filters []string
	argIndex := 1

	query := `SELECT id, months_duration, present_month, total_amount, amount_paid, last_payment_date, is_fully_paid, created_at FROM installment`

	if in.MonthsDuration > 0 {
		filters = append(filters, fmt.Sprintf("months_duration = $%d", argIndex))
		args = append(args, in.MonthsDuration)
		argIndex++
	}
	if in.CreatedAfter != "" {
		filters = append(filters, fmt.Sprintf("created_at >= $%d", argIndex))
		args = append(args, in.CreatedAfter)
		argIndex++
	}
	if in.CreatedBefore != "" {
		filters = append(filters, fmt.Sprintf("created_at <= $%d", argIndex))
		args = append(args, in.CreatedBefore)
		argIndex++
	}
	if in.TotalAmountMin > 0 {
		filters = append(filters, fmt.Sprintf("total_amount >= $%d", argIndex))
		args = append(args, in.TotalAmountMin)
		argIndex++
	}
	if in.TotalAmountMax > 0 {
		filters = append(filters, fmt.Sprintf("total_amount <= $%d", argIndex))
		args = append(args, in.TotalAmountMax)
		argIndex++
	}

	if len(filters) > 0 {
		query += " WHERE " + strings.Join(filters, " AND ")
	}

	rows, err := d.db.Queryx(query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to query installments: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var debt pb.Debt
		if err := rows.Scan(
			&debt.Id,
			&debt.MonthsDuration,
			&debt.PresentMonth,
			&debt.TotalAmount,
			&debt.AmountPaid,
			&debt.LastPaymentDate,
			&debt.IsFullyPaid,
			&debt.CreatedAt,
		); err != nil {
			return nil, fmt.Errorf("failed to scan installment: %w", err)
		}
		debts = append(debts, &debt)
	}

	return &pb.DebtsList{Debts: debts}, nil
}

// GetClientDebts retrieves all installments for a specific client
func (d *installmentRepo) GetClientDebts(in *pb.ClientID) (*pb.DebtsList, error) {
	var debts []*pb.Debt

	query := `
        SELECT id, months_duration, present_month, total_amount, amount_paid, last_payment_date, is_fully_paid, created_at
        FROM installment
        WHERE client_id = $1`

	rows, err := d.db.Queryx(query, in.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to get client installments: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var debt pb.Debt
		if err := rows.Scan(
			&debt.Id,
			&debt.MonthsDuration,
			&debt.PresentMonth,
			&debt.TotalAmount,
			&debt.AmountPaid,
			&debt.LastPaymentDate,
			&debt.IsFullyPaid,
			&debt.CreatedAt,
		); err != nil {
			return nil, fmt.Errorf("failed to scan installment: %w", err)
		}
		debts = append(debts, &debt)
	}

	return &pb.DebtsList{Debts: debts}, nil
}
