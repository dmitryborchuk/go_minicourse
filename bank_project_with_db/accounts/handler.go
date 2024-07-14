package accounts

import (
	"awesomeProject/accounts/dto"
	"awesomeProject/accounts/models"
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
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
	var request dto.CreateAccountRequest
	if err := c.Bind(&request); err != nil {
		c.Logger().Error(err)

		return c.String(http.StatusBadRequest, "invalid request")
	}

	if len(request.Name) == 0 {
		return c.String(http.StatusBadRequest, "empty name")
	}

	connectionString := "host=0.0.0.0 port=5432 dbname=postgres user=postgres password=mysecretpassword"
	db, err := sql.Open("pgx", connectionString)
	if err != nil {
		return fmt.Errorf("failed to connect to database")
	}

	defer func() {
		db.Close()
	}()

	if err := db.Ping(); err != nil {
		return fmt.Errorf("failed to ping the database")
	}

	_, err = db.Exec("INSERT INTO accounts(name, balance) VALUES($1, $2)", request.Name, request.Amount)

	if err != nil {
		return c.String(http.StatusForbidden, "account already exists")
	}

	return c.NoContent(http.StatusCreated)
}

func (h *Handler) GetAccount(c echo.Context) error {
	name := c.QueryParams().Get("name")

	connectionString := "host=0.0.0.0 port=5432 dbname=postgres user=postgres password=mysecretpassword"
	db, err := sql.Open("pgx", connectionString)
	if err != nil {
		return fmt.Errorf("failed to connect to database")
	}

	defer func() {
		db.Close()
	}()

	if err := db.Ping(); err != nil {
		return fmt.Errorf("failed to ping the database")
	}

	row := db.QueryRow("SELECT name, balance FROM accounts WHERE name = $1", name)

	var response dto.GetAccountResponse

	row.Scan(&response.Name, &response.Amount)

	if response.Name != name {
		return c.String(http.StatusNotFound, "account not found")
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

	connectionString := "host=0.0.0.0 port=5432 dbname=postgres user=postgres password=mysecretpassword"
	db, err := sql.Open("pgx", connectionString)
	if err != nil {
		return fmt.Errorf("failed to connect to database")
	}

	defer func() {
		db.Close()
	}()

	if err := db.Ping(); err != nil {
		return fmt.Errorf("failed to ping the database")
	}

	row := db.QueryRow("SELECT name, balance FROM accounts WHERE name = $1", request.Name)

	var response dto.GetAccountResponse

	row.Scan(&response.Name, &response.Amount)

	if response.Name != request.Name {
		return c.String(http.StatusNotFound, "account not found")
	}

	_, err = db.Exec("DELETE FROM accounts WHERE name = $1", request.Name)

	if row.Err() != nil {
		return c.String(http.StatusBadRequest, "account delete failed")
	}

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

	connectionString := "host=0.0.0.0 port=5432 dbname=postgres user=postgres password=mysecretpassword"
	db, err := sql.Open("pgx", connectionString)
	if err != nil {
		return fmt.Errorf("failed to connect to database")
	}

	defer func() {
		db.Close()
	}()

	if err := db.Ping(); err != nil {
		return fmt.Errorf("failed to ping the database")
	}

	row := db.QueryRow("SELECT name, balance FROM accounts WHERE name = $1", request.Name)

	var response dto.GetAccountResponse

	row.Scan(&response.Name, &response.Amount)

	if response.Name != request.Name {
		return c.String(http.StatusNotFound, "account not found")
	}

	_, err = db.Exec("UPDATE accounts SET balance=$2 WHERE name = $1", request.Name, request.Amount)

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

	connectionString := "host=0.0.0.0 port=5432 dbname=postgres user=postgres password=mysecretpassword"
	db, err := sql.Open("pgx", connectionString)
	if err != nil {
		return fmt.Errorf("failed to connect to database")
	}

	defer func() {
		db.Close()
	}()

	if err := db.Ping(); err != nil {
		return fmt.Errorf("failed to ping the database")
	}

	row := db.QueryRow("SELECT name, balance FROM accounts WHERE name = $1", request.Name)

	var response dto.GetAccountResponse

	row.Scan(&response.Name, &response.Amount)

	if response.Name != request.Name {
		return c.String(http.StatusNotFound, "account not found")
	}

	_, err = db.Exec("UPDATE accounts SET name=$2 WHERE name = $1", request.Name, request.NewName)

	return c.NoContent(http.StatusOK)
}
