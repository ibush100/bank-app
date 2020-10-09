package transaction

type UpdateUserBalance struct {
	Username string
	Password string
	Email    string
	TopUp    int
}

type Transaction struct {
	PayeeEmail string `validate:"required"`
	PayorEmail string `validate:"required"`
	Amount     int    `validate:"required"`
}
