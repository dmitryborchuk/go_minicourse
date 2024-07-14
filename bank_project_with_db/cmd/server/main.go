package main

import (
	"awesomeProject/accounts"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"strconv"
)

func main() {
	accountsHandler := accounts.New()

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/account", accountsHandler.GetAccount)
	e.POST("/account/create", accountsHandler.CreateAccount)
	e.POST("/account/delete", accountsHandler.DeleteAccount)
	e.POST("/account/change/balance", accountsHandler.ChangeAccountsBalance)
	e.POST("/account/change/name", accountsHandler.ChangeAccountsName)

	// Start server
	var portNumber int
	fmt.Printf("Choose the port number: ")
	fmt.Scanln(&portNumber)

	if portNumber < 0 || portNumber > 65535 {
		panic("port number must be between 0 and 65535")
	}

	e.Logger.Fatal(e.Start(":" + strconv.Itoa(portNumber)))
}
