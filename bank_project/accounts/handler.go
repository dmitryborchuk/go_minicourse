package accounts

import (
	"awesomeProject/accounts/dto"
	"awesomeProject/accounts/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"sync"
)

func New() *Handler {
	return &Handler{
		accounts: make(map[string]*models.Account),
		guard:    &sync.RWMutex{},
	}
}

type Handler struct {
	accounts map[string]*models.Account
	guard    *sync.RWMutex
}

func (h *Handler) CreateAccount(c echo.Context) error {
	var request dto.CreateAccountRequest // {"name": "alice", "amount": 50}
	if err := c.Bind(&request); err != nil {
		c.Logger().Error(err)

		return c.String(http.StatusBadRequest, "invalid request")
	}

	if len(request.Name) == 0 {
		return c.String(http.StatusBadRequest, "empty name")
	}

	h.guard.Lock()

	if _, ok := h.accounts[request.Name]; ok {
		h.guard.Unlock()

		return c.String(http.StatusForbidden, "account already exists")
	}

	h.accounts[request.Name] = &models.Account{
		Name:   request.Name,
		Amount: request.Amount,
	}

	h.guard.Unlock()

	return c.NoContent(http.StatusCreated)
}

func (h *Handler) GetAccount(c echo.Context) error {
	name := c.QueryParams().Get("name")

	h.guard.RLock()

	account, ok := h.accounts[name]

	h.guard.RUnlock()

	if !ok {
		return c.String(http.StatusNotFound, "account not found")
	}

	response := dto.GetAccountResponse{
		Name:   account.Name,
		Amount: account.Amount,
	}

	return c.JSON(http.StatusOK, response)
}

func (h *Handler) DeleteAccount(c echo.Context) error {
	var request dto.DeleteAccountRequest
	if err := c.Bind(&request); err != nil {
		c.Logger().Error(err)

		return c.String(http.StatusBadRequest, "invalid request")
	}

	if len(request.Name) == 0 {
		return c.String(http.StatusBadRequest, "empty name")
	}

	h.guard.Lock()

	if _, ok := h.accounts[request.Name]; !ok {
		h.guard.Unlock()

		return c.String(http.StatusForbidden, "account does not exist")
	}

	delete(h.accounts, request.Name)

	h.guard.Unlock()

	return c.NoContent(http.StatusOK)
}

func (h *Handler) ChangeAccountsBalance(c echo.Context) error {
	var request dto.ChangeAccountsBalanceRequest
	if err := c.Bind(&request); err != nil {
		c.Logger().Error(err)

		return c.String(http.StatusBadRequest, "invalid request")
	}

	if len(request.Name) == 0 {
		return c.String(http.StatusBadRequest, "empty name")
	}

	h.guard.Lock()

	if _, ok := h.accounts[request.Name]; !ok {
		h.guard.Unlock()

		return c.String(http.StatusForbidden, "account does not exist")
	}

	h.accounts[request.Name].Amount = request.Amount

	h.guard.Unlock()

	return c.NoContent(http.StatusOK)
}

func (h *Handler) ChangeAccountsName(c echo.Context) error {
	var request dto.ChangeAccountsNameRequest
	if err := c.Bind(&request); err != nil {
		c.Logger().Error(err)

		return c.String(http.StatusBadRequest, "invalid request")
	}

	if len(request.Name) == 0 {
		return c.String(http.StatusBadRequest, "empty name")
	}

	if len(request.NewName) == 0 {
		return c.String(http.StatusBadRequest, "empty new_name")
	}

	h.guard.Lock()

	bal := h.accounts[request.Name].Amount
	delete(h.accounts, request.Name)

	h.accounts[request.NewName] = &models.Account{
		Name:   request.NewName,
		Amount: bal,
	}

	h.guard.Unlock()

	return c.NoContent(http.StatusOK)
}
