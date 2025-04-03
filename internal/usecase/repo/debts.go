package repo

import (
	"context"
	"database/sql"
	"debts-service/internal/usecase/entity"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

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

func (d *installmentRepo) CreateDebt(in *pb.DebtsRequest) (*pb.Debts, error) {
	var debt pb.Debts
	var lastPaymentDate, saleID, shouldPayAt sql.NullString

	var saleIDParam interface{}
	if in.SaleId == "" {
		saleIDParam = nil
	} else {
		saleIDParam = in.SaleId
	}

	var shouldPayAtParam interface{}
	if in.ShouldPayAt == "" {
		shouldPayAtParam = nil
	} else {
		parsedDate, err := time.Parse("02-01-2006", in.ShouldPayAt)
		if err != nil {
			return nil, fmt.Errorf("invalid date format for should_pay_at, expected dd-mm-yyyy: %w", err)
		}
		shouldPayAtParam = parsedDate.Format("2006-01-02")
	}

	query := `
		INSERT INTO installment (id, client_id, sale_id, total_amount, amount_paid,
		                         currency_code, should_pay_at, debt_type, company_id)
		VALUES (gen_random_uuid(), $1, $2, $3, 0, $4, $5, $6, $7)
		RETURNING id, client_id, sale_id, total_amount, amount_paid, last_payment_date, is_fully_paid,
		          currency_code, should_pay_at, debt_type, created_at, company_id`

	err := d.db.QueryRowx(query,
		in.ClientId,
		saleIDParam,
		in.TotalAmount,
		in.CurrencyCode,
		shouldPayAtParam,
		in.DebtType,
		in.CompanyId,
	).Scan(
		&debt.Id,
		&debt.ClientId,
		&saleID,
		&debt.TotalAmount,
		&debt.AmountPaid,
		&lastPaymentDate,
		&debt.IsFullyPaid,
		&debt.CurrencyCode,
		&shouldPayAt,
		&debt.DebtType,
		&debt.CreatedAt,
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

	if saleID.Valid {
		debt.SaleId = saleID.String
	} else {
		debt.SaleId = ""
	}

	if shouldPayAt.Valid {
		debt.ShouldPayAt = shouldPayAt.String
	} else {
		debt.ShouldPayAt = ""
	}

	return &debt, nil
}

// GetDebt retrieves an installment by ID
func (d *installmentRepo) GetDebt(in *pb.DebtsID) (*pb.Debts, error) {
	var debt pb.Debts
	var lastPaymentDate, saleID, shouldPayAt sql.NullString

	query := `
        SELECT id, client_id, sale_id, total_amount, amount_paid, last_payment_date, is_fully_paid,
		    currency_code, should_pay_at, debt_type, created_at, company_id
        FROM installment
        WHERE id = $1 AND company_id = $2`

	err := d.db.QueryRowx(query, in.Id, in.CompanyId).Scan(
		&debt.Id,
		&debt.ClientId,
		&saleID,
		&debt.TotalAmount,
		&debt.AmountPaid,
		&lastPaymentDate,
		&debt.IsFullyPaid,
		&debt.CurrencyCode,
		&shouldPayAt,
		&debt.DebtType,
		&debt.CreatedAt,
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

	if saleID.Valid {
		debt.SaleId = saleID.String
	} else {
		debt.SaleId = ""
	}

	if shouldPayAt.Valid {
		debt.ShouldPayAt = shouldPayAt.String
	} else {
		debt.ShouldPayAt = ""
	}

	debt.BalanceOfDebt = debt.TotalAmount - debt.AmountPaid

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
		argIdx  = 3
	)

	query := `
        SELECT 
            id, 
            client_id, 
            sale_id,
            total_amount, 
            amount_paid, 
            last_payment_date, 
            is_fully_paid, 
            created_at, 
            currency_code, 
            debt_type,
            should_pay_at,
            company_id,
            COUNT(*) OVER() AS total_count
        FROM installment 
        WHERE company_id = $1 and debt_type = $2
    `
	args = append(args, in.CompanyId, in.DebtType)

	if in.IsFullyPay != "" {
		boolVal, err := strconv.ParseBool(in.IsFullyPay)
		if err != nil {
			return nil, fmt.Errorf("invalid is_fully_pay value: %w", err)
		}
		filters = append(filters, fmt.Sprintf("is_fully_paid = $%d", argIdx))
		args = append(args, boolVal)
		argIdx++
	}

	if in.CurrencyCode != "" {
		filters = append(filters, fmt.Sprintf("currency_code = $%d", argIdx))
		args = append(args, in.CurrencyCode)
		argIdx++
	}

	if in.NoPaidDebt {
		filters = append(filters, fmt.Sprintf("amount_paid = $%d", argIdx))
		args = append(args, 0)
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
		argIdx++
	}

	rows, err := d.db.Queryx(query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to query installments: %w", err)
	}
	defer rows.Close()

	var totalCount int64
	for rows.Next() {
		var debt pb.Debts
		var lastPaymentDate, shouldPayAt sql.NullTime
		var saleID sql.NullString

		if err := rows.Scan(
			&debt.Id,
			&debt.ClientId,
			&saleID,
			&debt.TotalAmount,
			&debt.AmountPaid,
			&lastPaymentDate,
			&debt.IsFullyPaid,
			&debt.CreatedAt,
			&debt.CurrencyCode,
			&debt.DebtType,
			&shouldPayAt,
			&debt.CompanyId,
			&totalCount,
		); err != nil {
			return nil, fmt.Errorf("failed to scan installment: %w", err)
		}

		if lastPaymentDate.Valid {
			debt.LastPaymentDate = lastPaymentDate.Time.Format("2006-01-02")
		} else {
			debt.LastPaymentDate = ""
		}

		if shouldPayAt.Valid {
			debt.ShouldPayAt = shouldPayAt.Time.Format("2006-01-02")
		} else {
			debt.ShouldPayAt = ""
		}

		if saleID.Valid {
			debt.SaleId = saleID.String
		} else {
			debt.ShouldPayAt = ""
		}

		debt.BalanceOfDebt = debt.TotalAmount - debt.AmountPaid

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

func (d *installmentRepo) GetClientDebts(in *pb.ClientID) (*pb.DebtsList, error) {
	if in.Id == "" {
		return nil, fmt.Errorf("client id is required")
	}

	var debts []*pb.Debts

	query := `
        SELECT 
            id, 
            client_id, 
            sale_id,
            total_amount, 
            amount_paid, 
            last_payment_date, 
            is_fully_paid, 
            created_at, 
            currency_code,
            debt_type,
            should_pay_at,
            company_id
        FROM installment
        WHERE client_id = $1 AND company_id = $2 AND debt_type = $3
    `

	rows, err := d.db.Queryx(query, in.Id, in.CompanyId, in.DebtType)
	if err != nil {
		return nil, fmt.Errorf("failed to get client installments: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var debt pb.Debts
		var lastPaymentDate sql.NullTime
		var shouldPayAt sql.NullTime
		var saleID sql.NullString

		if err := rows.Scan(
			&debt.Id,
			&debt.ClientId,
			&saleID,
			&debt.TotalAmount,
			&debt.AmountPaid,
			&lastPaymentDate,
			&debt.IsFullyPaid,
			&debt.CreatedAt,
			&debt.CurrencyCode,
			&debt.DebtType,
			&shouldPayAt,
			&debt.CompanyId,
		); err != nil {
			return nil, fmt.Errorf("failed to scan installment: %w", err)
		}

		if lastPaymentDate.Valid {
			debt.LastPaymentDate = lastPaymentDate.Time.Format("2006-01-02")
		} else {
			debt.LastPaymentDate = ""
		}

		if shouldPayAt.Valid {
			debt.ShouldPayAt = shouldPayAt.Time.Format("2006-01-02")
		} else {
			debt.ShouldPayAt = ""
		}

		if saleID.Valid {
			debt.SaleId = saleID.String
		} else {
			debt.SaleId = ""
		}

		debt.BalanceOfDebt = debt.TotalAmount - debt.AmountPaid

		debts = append(debts, &debt)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return &pb.DebtsList{Installments: debts}, nil
}

func (d *installmentRepo) PayPayment(in *pb.PayDebtsReq) (*pb.Debts, error) {
	paymentID := uuid.NewString()

	tx, err := d.db.Beginx()
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	paymentQuery := `
        INSERT INTO payments (id, installment_id, payment_amount, pay_type, payment_date)
        VALUES ($1, $2, $3, $4, CURRENT_TIMESTAMP)
    `
	_, err = tx.Exec(paymentQuery, paymentID, in.DebtId, in.PaidAmount, in.PayType)
	if err != nil {
		return nil, fmt.Errorf("failed to insert payment: %w", err)
	}

	debtQuery := `
        UPDATE installment
        SET amount_paid = amount_paid + $1,
            last_payment_date = CURRENT_TIMESTAMP
        WHERE id = $2 AND company_id = $3
        RETURNING 
            id, 
            client_id, 
            sale_id, 
            total_amount, 
            amount_paid, 
            last_payment_date, 
            is_fully_paid, 
            created_at, 
            currency_code,
            debt_type,
            should_pay_at,
            company_id
    `
	var debt pb.Debts
	var lastPaymentDate sql.NullTime
	var shouldPayAt sql.NullTime
	var saleID sql.NullString

	err = tx.QueryRowx(debtQuery, in.PaidAmount, in.DebtId, in.CompanyId).Scan(
		&debt.Id,
		&debt.ClientId,
		&saleID,
		&debt.TotalAmount,
		&debt.AmountPaid,
		&lastPaymentDate,
		&debt.IsFullyPaid,
		&debt.CreatedAt,
		&debt.CurrencyCode,
		&debt.DebtType,
		&shouldPayAt,
		&debt.CompanyId,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("debt not found or company_id mismatch: %w", err)
		}
		return nil, fmt.Errorf("failed to update installment: %w", err)
	}

	if lastPaymentDate.Valid {
		debt.LastPaymentDate = lastPaymentDate.Time.Format("2006-01-02")
	} else {
		debt.LastPaymentDate = ""
	}

	if shouldPayAt.Valid {
		debt.ShouldPayAt = shouldPayAt.Time.Format("2006-01-02")
	} else {
		debt.ShouldPayAt = ""
	}

	if saleID.Valid {
		debt.SaleId = saleID.String
	} else {
		debt.SaleId = ""
	}

	debt.BalanceOfDebt = debt.TotalAmount - debt.AmountPaid

	if err = tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return &debt, nil
}

func (d *installmentRepo) GetTotalDebtSum(in *pb.CompanyID) (*pb.SumMoney, error) {
	query := `
		SELECT currency_code, SUM(GREATEST(total_amount - amount_paid, 0)) AS debt
		FROM installment
		WHERE company_id = $1 AND debt_type = $2 
		GROUP BY currency_code`
	args := []interface{}{in.Id, in.DebtType} // Добавлен второй параметр

	rows, err := d.db.QueryContext(context.Background(), query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query for company debt: %w", err)
	}
	defer rows.Close()

	result := &pb.SumMoney{
		CompanyId: in.Id,
		Sum:       make([]*pb.Money, 0),
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

// GetUserTotalDebtSum рассчитывает сумму оставшихся долгов для конкретного клиента,
func (d *installmentRepo) GetUserTotalDebtSum(in *pb.ClientID) (*pb.SumMoney, error) {
	query := `
		SELECT currency_code, SUM(GREATEST(total_amount - amount_paid, 0)) AS debt
		FROM installment
		WHERE client_id = $1 AND debt_type = $2 AND company_id = $3
		GROUP BY currency_code
	`
	rows, err := d.db.QueryContext(context.Background(), query, in.Id, in.DebtType, in.CompanyId)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query for client debt: %w", err)
	}
	defer rows.Close()

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

func (d *installmentRepo) GetDebtsForExel(in *pb.FilterExelDebt) (*entity.ListDebtsExelDb, error) {
	query := `
		SELECT 
			client_id,
			COALESCE(
				json_agg(
					json_build_object(
						'total_amount', total_amount,
						'amount_paid', amount_paid,
						'debt_balance', total_amount - amount_paid,
						'last_paid_date', COALESCE(last_payment_date::text, ''),
						'currency_code', currency_code
					)
				), '[]'
			) AS debts
		FROM installment
		WHERE company_id = $1 and is_fully_paid = false
		GROUP BY client_id
	`

	var tmpResults []struct {
		ClientID string `db:"client_id"`
		Debts    []byte `db:"debts"`
	}
	if err := d.db.Select(&tmpResults, query, in.CompanyId); err != nil {
		return nil, fmt.Errorf("query execution failed: %w", err)
	}

	listDebts := &entity.ListDebtsExelDb{
		Debts: make([]*entity.DebtsExelDbRes, len(tmpResults)),
	}

	for i, row := range tmpResults {
		var userDebts []*entity.UserDebts
		if err := json.Unmarshal(row.Debts, &userDebts); err != nil {
			return nil, fmt.Errorf("failed to unmarshal debts JSON: %w", err)
		}

		listDebts.Debts[i] = &entity.DebtsExelDbRes{
			ClientID: row.ClientID,
			Debts:    userDebts,
		}
	}

	return listDebts, nil
}
