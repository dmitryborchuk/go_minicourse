package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc/proto"
	"time"
)

func main() {
	cmdVal := flag.String("cmd", "", "command to execute")
	nameVal := flag.String("name", "", "name of account")
	newNameVal := flag.String("new_name", "", "new name of account")
	amountVal := flag.Int("amount", 0, "amount of account")

	flag.Parse()

	conn, err := grpc.NewClient("0.0.0.0:1232", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	c := proto.NewBankClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	defer func() {
		_ = conn.Close()
	}()

	switch *cmdVal {
	case "create":
		_, err := c.CreateAccount(ctx, &proto.CreateAccountRequest{Name: *nameVal, Amount: (int32)(*amountVal)})
		if err != nil {
			panic(err)
		}
	case "get":
		resp, err := c.GetAccount(ctx, &proto.GetAccountRequest{Name: *nameVal})
		if err != nil {
			panic(err)
		}
		fmt.Printf("response account name: %s and amount: %d\n", resp.Name, resp.Amount)
	case "delete":
		_, err := c.DeleteAccount(ctx, &proto.DeleteAccountRequest{Name: *nameVal})
		if err != nil {
			panic(err)
		}
	case "change_name":
		_, err := c.ChangeAccountsName(ctx, &proto.ChangeAccountsNameRequest{Name: *nameVal, NewName: *newNameVal})
		if err != nil {
			panic(err)
		}
	case "change_balance":
		_, err := c.ChangeAccountsBalance(ctx, &proto.ChangeAccountsBalanceRequest{Name: *nameVal, Amount: (int32)(*amountVal)})
		if err != nil {
			panic(err)
		}
	default:
		panic(fmt.Errorf("Invalid command given"))
	}
}
