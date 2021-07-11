package bank_balance_history

import "time"

type BankBalanceHistory struct {
	ID             	int `gorm:"primaryKey"`
	BankBalanceId  	int
	BalanceBefore	int
	BalanceAfter	int
	Activity		string
	Type			string
	Ip				string
	Location		string
	UserAgent		string
	Author			string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
