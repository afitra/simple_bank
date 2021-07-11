package bank_balance

import "time"

type BankBalance struct {
	ID             	int `gorm:"primaryKey"`
	UserId			int
	Balance			string
	BalanceAchieve	int
	Code			string
	Enable			bool
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
