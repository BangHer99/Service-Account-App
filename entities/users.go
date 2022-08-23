package entities

type Users struct {
	NoTelp    int
	Password  string
	Name      string
	Gender    string
	Balance   int
	Currency  string
	CreatedAt string
	UpdateAt  string
}

type OtherUser struct {
	NoTelp     int
	Name       string
	Gender     string
	Created_at string
	Updated_at string
}
