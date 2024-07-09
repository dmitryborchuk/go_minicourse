package dto

type CreateAccountRequest struct {
	Name   string `json:"name"`
	Amount int    `json:"amount"`
}

type GetAccountRequest struct {
	Name string `json:"name"`
}

type ChangeAccountsBalanceRequest struct {
	Name   string `json:"name"`
	Amount int    `json:"amount"`
}

type ChangeAccountsNameRequest struct {
	Name    string `json:"name"`
	NewName string `json:"new_name"`
}

type DeleteAccountRequest struct {
	Name string `json:"name"`
}
