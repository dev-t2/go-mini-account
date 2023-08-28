package accounts

type account struct {
	id      int
	balance int
}

func CreateAccount(id int) *account {
	return &account{id: id, balance: 0}
}