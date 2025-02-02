package repo

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	pb "debts-service/internal/generated/debts"
	"debts-service/internal/usecase"
	"github.com/google/uuid"

	"github.com/jmoiron/sqlx"
)

type installmentRepo struct {
	db *sqlx.DB
}

func NewInstallmentRepo(db *sqlx.DB) usecase.DebtsRepo {
	return &installmentRepo{db: db}
}

// CreateDebt creates a new installment record
func (d *installmentRepo) CreateDebt(in *pb.DebtsRequest) (*pb.Debts, error) {
	var debt pb.Debts
	var lastPaymentDate sql.NullString
	query := `
		INSERT INTO installment (id, client_id, total_amount, amount_paid, currency_code, company_id)
		VALUES (gen_random_uuid(), $1, $2, 0, $3, $4)
		RETURNING id, client_id, total_amount, amount_paid, last_payment_date, is_fully_paid, created_at, currency_code, company_id`

	err := d.db.QueryRowx(query, in.ClientId, in.TotalAmount, in.CurrencyCode, in.CompanyId).Scan(
		&debt.Id,
		&debt.ClientId,
		&debt.TotalAmount,
		&debt.AmountPaid,
		&lastPaymentDate,
		&debt.IsFullyPaid,
		&debt.CreatedAt,
		&debt.CurrencyCode,
		&debt.CompanyId,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create installment: %w", err)
	}
	if lastPaymentDate.Valid {
		debt.LastPaymentDate = lastPaymentDate.String
	} else {
		debt.LastPaymentDate = ""
	}
	return &debt, nil
}

// GetDebt retrieves an installment by ID
func (d *installmentRepo) GetDebt(in *pb.DebtsID) (*pb.Debts, error) {
	var debt pb.Debts
	var lastPaymentDate sql.NullString

	query := `
        SELECT id, client_id, total_amount, amount_paid, last_payment_date, is_fully_paid, created_at, currency_code, company_id
        FROM installment
        WHERE id = $1 AND company_id = $2`

	err := d.db.QueryRowx(query, in.Id, in.CompanyId).Scan(
		&debt.Id,
		&debt.ClientId,
		&debt.TotalAmount,
		&debt.AmountPaid,
		&lastPaymentDate,
		&debt.IsFullyPaid,
		&debt.CreatedAt,
		&debt.CurrencyCode,
		&debt.CompanyId,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("installment not found")
		}
		return nil, fmt.Errorf("failed to get installment: %w", err)
	}
	if lastPaymentDate.Valid {
		debt.LastPaymentDate = lastPaymentDate.String
	} else {
		debt.LastPaymentDate = ""
	}

	return &debt, nil
}

func (d *installmentRepo) GetListDebts(in *pb.FilterDebts) (*pb.DebtsList, error) {
	if in.CompanyId == "" {
		return nil, fmt.Errorf("company_id is required")
	}

	var (
		debts   []*pb.Debts
		args    []interface{}
		filters []string
		argIdx  = 2 // $1 уже используется для company_id
	)

	query := `
        SELECT 
            id, 
            client_id, 
            total_amount, 
            amount_paid, 
            last_payment_date, 
            is_fully_paid, 
            created_at, 
            currency_code, 
            company_id,
            COUNT(*) OVER() AS total_count
        FROM installment 
        WHERE company_id = $1
    `
	args = append(args, in.CompanyId)

	if in.IsFullyPay != "" {
		filters = append(filters, fmt.Sprintf("is_fully_paid = $%d", argIdx))
		args = append(args, in.IsFullyPay)
		argIdx++
	}

	if in.TotalAmountMin > 0 {
		filters = append(filters, fmt.Sprintf("total_amount >= $%d", argIdx))
		args = append(args, in.TotalAmountMin)
		argIdx++
	}
	if in.TotalAmountMax > 0 {
		filters = append(filters, fmt.Sprintf("total_amount <= $%d", argIdx))
		args = append(args, in.TotalAmountMax)
		argIdx++
	}

	// Фильтр по валюте
	if in.CurrencyCode != "" {
		filters = append(filters, fmt.Sprintf("currency_code = $%d", argIdx))
		args = append(args, in.CurrencyCode)
		argIdx++
	}

	if in.Description != "" {
		filters = append(filters, fmt.Sprintf("description ILIKE $%d", argIdx))
		args = append(args, "%"+in.Description+"%")
		argIdx++
	}

	if len(filters) > 0 {
		query += " AND " + strings.Join(filters, " AND ")
	}

	query += " ORDER BY created_at DESC"

	if in.Limit > 0 && in.Page > 0 {

		query += fmt.Sprintf(" LIMIT $%d", argIdx)
		args = append(args, int64(in.Limit))
		argIdx++

		query += fmt.Sprintf(" OFFSET $%d", argIdx)
		offset := int64(in.Limit * (in.Page - 1))
		args = append(args, offset)
	}

	rows, err := d.db.Queryx(query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to query installments: %w", err)
	}
	defer rows.Close()

	var totalCount int64
	for rows.Next() {
		var debt pb.Debts
		var lastPaymentDate sql.NullTime

		if err := rows.Scan(
			&debt.Id,
			&debt.ClientId,
			&debt.TotalAmount,
			&debt.AmountPaid,
			&lastPaymentDate,
			&debt.IsFullyPaid,
			&debt.CreatedAt,
			&debt.CurrencyCode,
			&debt.CompanyId,
			&totalCount,
		); err != nil {
			return nil, fmt.Errorf("failed to scan installment: %w", err)
		}

		if lastPaymentDate.Valid {
			// Например, формат "YYYY-MM-DD" (можно изменить по необходимости)
			debt.LastPaymentDate = lastPaymentDate.Time.Format("2006-01-02")
		} else {
			debt.LastPaymentDate = ""
		}

		debts = append(debts, &debt)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return &pb.DebtsList{
		Installments: debts,
		TotalCount:   totalCount,
	}, nil
}

// GetClientDebts retrieves all installments for a specific client
func (d *installmentRepo) GetClientDebts(in *pb.ClientID) (*pb.DebtsList, error) {
	var debts []*pb.Debts

	query := `
        SELECT id, client_id, total_amount, amount_paid, last_payment_date, is_fully_paid, created_at, currency_code, company_id
        FROM installment
        WHERE client_id = $1`

	rows, err := d.db.Queryx(query, in.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to get client installments: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var debt pb.Debts
		var lastPaymentDate sql.NullString
		if err := rows.Scan(
			&debt.Id,
			&debt.ClientId,
			&debt.TotalAmount,
			&debt.AmountPaid,
			&lastPaymentDate,
			&debt.IsFullyPaid,
			&debt.CreatedAt,
			&debt.CurrencyCode,
			&debt.CompanyId,
		); err != nil {
			return nil, fmt.Errorf("failed to scan installment: %w", err)
		}
		if lastPaymentDate.Valid {
			debt.LastPaymentDate = lastPaymentDate.String
		} else {
			debt.LastPaymentDate = ""
		}
		debts = append(debts, &debt)
	}

	return &pb.DebtsList{Installments: debts}, nil
}

// PayPayment records a payment for an installment
func (d *installmentRepo) PayPayment(in *pb.PayDebtsReq) (*pb.Debts, error) {
	// Проверяем, что DebtId не пустой
	if in.DebtId == "" {
		return nil, fmt.Errorf("invalid input: DebtId is required")
	}

	id := uuid.NewString()

	tx, err := d.db.Beginx()
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	// Insert payment record
	paymentQuery := `
        INSERT INTO payments (id, installment_id, payment_amount, payment_date)
        VALUES ($1, $2, $3, CURRENT_TIMESTAMP)`

	_, err = tx.Exec(paymentQuery, id, in.DebtId, in.PaidAmount)
	if err != nil {
		return nil, fmt.Errorf("failed to insert payment: %w", err)
	}

	// Update installment record (do not modify is_fully_paid column directly)
	var debt pb.Debts
	debtQuery := `
        UPDATE installment
        SET amount_paid = amount_paid + $1,
            last_payment_date = CURRENT_TIMESTAMP
        WHERE id = $2
        RETURNING id, client_id, total_amount, amount_paid, last_payment_date, is_fully_paid, created_at, currency_code, company_id`

	err = tx.QueryRowx(debtQuery, in.PaidAmount, in.DebtId).Scan(
		&debt.Id,
		&debt.ClientId,
		&debt.TotalAmount,
		&debt.AmountPaid,
		&debt.LastPaymentDate,
		&debt.IsFullyPaid,
		&debt.CreatedAt,
		&debt.CurrencyCode,
		&debt.CompanyId,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("debt not found: %w", err)
		}
		return nil, fmt.Errorf("failed to update installment: %w", err)
	}

	// Commit transaction
	if err = tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return &debt, nil
}

func (d *installmentRepo) GetTotalDebtSum(in *pb.CompanyID) (*pb.SumMoney, error) {

	query := `
		SELECT currency_code, SUM(GREATEST(total_amount - amount_paid, 0)) AS debt
		FROM installment
		WHERE company_id = $1
		GROUP BY currency_code
	`
	rows, err := d.db.QueryContext(context.Background(), query, in.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query for company debt: %w", err)
	}
	defer rows.Close()

	// Формируем ответный объект SumMoney.
	result := &pb.SumMoney{
		CompanyId: in.Id,
		Sum:       make([]*pb.Money, 0),
	}

	for rows.Next() {
		var currency string
		var debt float64 // DECIMAL(10,2) можно сканировать в float64
		if err := rows.Scan(&currency, &debt); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		result.Sum = append(result.Sum, &pb.Money{
			Currency: currency,
			Sum:      debt,
		})
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}
	return result, nil
}

// GetUserTotalDebtSum рассчитывает сумму оставшихся долгов для конкретного клиента,
// также группируя задолженности по валюте.
func (d *installmentRepo) GetUserTotalDebtSum(in *pb.ClientID) (*pb.SumMoney, error) {
	// SQL-запрос: аналогичный запрос, но фильтруем по client_id.
	query := `
		SELECT currency_code, SUM(GREATEST(total_amount - amount_paid, 0)) AS debt
		FROM installment
		WHERE client_id = $1
		GROUP BY currency_code
	`
	rows, err := d.db.QueryContext(context.Background(), query, in.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query for client debt: %w", err)
	}
	defer rows.Close()

	// Для ответа используем SumMoney; поле CompanyId можно оставить пустым.
	result := &pb.SumMoney{
		Sum: make([]*pb.Money, 0),
	}

	for rows.Next() {
		var currency string
		var debt float64
		if err := rows.Scan(&currency, &debt); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		result.Sum = append(result.Sum, &pb.Money{
			Currency: currency,
			Sum:      debt,
		})
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}
	return result, nil
}
