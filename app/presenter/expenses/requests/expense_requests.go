package requests

import (
	"time"

	"github.com/Rahmanwghazi/Monefy/business/expenses"
)

type Expense struct {
	UserID      uint      `json:"user_id"`
	Total       int       `json:"total"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (expense *Expense) ToDomain() expenses.ExpenseDomain {
	return expenses.ExpenseDomain{
		UserID:      expense.UserID,
		Total:       expense.Total,
		Description: expense.Description,
	}
}