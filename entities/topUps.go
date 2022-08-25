package entities

type TopUps struct {
	Id          int
	AccountTelp int
	Amount      int
	CreatedAt   string
}
type TopUpHistory struct {
	id              int
	NameUser        string
	To_account_name string
	Amount          int
	CreatedAt       string
}
