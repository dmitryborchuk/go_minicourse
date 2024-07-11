package accounts

import (
	"fmt"
	"grpc/accounts/models"
	"grpc/proto"
	"sync"
)

func New() *Handler {
	return &Handler{
		Accounts: make(map[string]*models.Account),
		Guard:    &sync.RWMutex{},
	}
}

type Handler struct {
	Accounts map[string]*models.Account
	Guard    *sync.RWMutex
}

func (h *Handler) CreateAccount(c *proto.CreateAccountRequest) error {
	if len(c.Name) == 0 {
		return fmt.Errorf("empty name")
	}

	h.Guard.Lock()

	if _, ok := h.Accounts[c.Name]; ok {
		h.Guard.Unlock()

		return fmt.Errorf("account already exists")
	}

	h.Accounts[c.Name] = &models.Account{
		Name:   c.Name,
		Amount: c.Amount,
	}

	h.Guard.Unlock()

	return nil
}

func (h *Handler) GetAccount(c *proto.GetAccountRequest) (*proto.GetAccountResponse, error) {
	name := c.Name

	h.Guard.RLock()

	account, ok := h.Accounts[name]

	h.Guard.RUnlock()

	if !ok {
		return nil, fmt.Errorf("account not found")
	}

	var response proto.GetAccountResponse
	response = proto.GetAccountResponse{
		Name:   account.Name,
		Amount: account.Amount,
	}

	return &response, nil
}

func (h *Handler) DeleteAccount(c *proto.DeleteAccountRequest) error {
	if len(c.Name) == 0 {
		return fmt.Errorf("empty name")
	}

	h.Guard.Lock()

	if _, ok := h.Accounts[c.Name]; !ok {
		h.Guard.Unlock()

		return fmt.Errorf("account does not exist")
	}

	delete(h.Accounts, c.Name)

	h.Guard.Unlock()

	return nil
}

func (h *Handler) ChangeAccountsBalance(c *proto.ChangeAccountsBalanceRequest) error {
	if len(c.Name) == 0 {
		return fmt.Errorf("empty name")
	}

	h.Guard.Lock()

	if _, ok := h.Accounts[c.Name]; !ok {
		h.Guard.Unlock()

		return fmt.Errorf("account does not exist")
	}

	h.Accounts[c.Name].Amount = c.Amount

	h.Guard.Unlock()

	return nil
}

func (h *Handler) ChangeAccountsName(c *proto.ChangeAccountsNameRequest) error {
	if len(c.Name) == 0 {
		return fmt.Errorf("empty name")
	}

	if len(c.Name) == 0 {
		return fmt.Errorf("empty new_name")
	}

	h.Guard.Lock()

	bal := h.Accounts[c.Name].Amount
	delete(h.Accounts, c.Name)

	h.Accounts[c.NewName] = &models.Account{
		Name:   c.NewName,
		Amount: bal,
	}

	h.Guard.Unlock()

	return nil
}
