package user_balance

 

type UserBalance struct {
	ID				int `gorm:"primaryKey"`
	UserId			int
	Balance			string
	BalanceAchieve	int
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
