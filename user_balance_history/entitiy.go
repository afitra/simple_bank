package user_balance_history

import "time"

type UserBalanceHistory struct {
	ID             int `gorm:"primaryKey"`
	UserBalanceId	int
	BalanceBefore	int
	BalanceAfter	int
	Activity		string
	Type 			string
	Ip				string
	Location		string
	UserAgent		string
	Author			string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}