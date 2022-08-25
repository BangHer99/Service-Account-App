package entities

import "time"

type TopUps struct {
	Id          int
	AccountTelp int
	Amount      int
	CreatedAt   time.Time
}

type TopUpHistory struct {
	Id              int
	NameUser        string
	To_account_name string
	Amount          int
	CreatedAt       string
}
