package entities

type Transfers struct {
	Id                int
	From_account_telp int
	To_account_telp   int
	Amount            int
	Created_at        string
}

type TransferHistory struct {
	Id                int
	From_account_name string
	To_account_name   string
	Amount            int
	Created_at        string
}
