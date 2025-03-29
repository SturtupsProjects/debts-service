package entity

type ListDebtsExelDb struct {
	Debts []*DebtsExelDbRes `db:"debts"`
}

type DebtsExelDbRes struct {
	ClientID string       `db:"client_id"`
	Debts    []*UserDebts `db:"debts"`
}

type UserDebts struct {
	TotalAmount  float64 `json:"total_amount"`
	AmountPaid   float64 `json:"amount_paid"`
	DebtBalance  float64 `json:"debt_balance"`
	LastPaidDate string  `json:"last_paid_date"`
	CurrencyCode string  `json:"currency_code"`
}
