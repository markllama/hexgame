package types

type User struct {
	Username string `json:"username"`
	FirstName string `json:"first_name"`
	Surname string `json:"surname"`
	EmailAddress string `json:"email_address"`
}
