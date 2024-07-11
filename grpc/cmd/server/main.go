package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"grpc/accounts"
	"grpc/proto"
	"net"
)

type server struct {
	proto.UnimplementedBankServer
	accountsHandler accounts.Handler
}

func (s *server) CreateAccount(ctx context.Context, req *proto.CreateAccountRequest) (*emptypb.Empty, error) {
	if err := s.accountsHandler.CreateAccount(req); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *server) GetAccount(ctx context.Context, req *proto.GetAccountRequest) (*proto.GetAccountResponse, error) {
	res, err := s.accountsHandler.GetAccount(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *server) DeleteAccount(ctx context.Context, req *proto.DeleteAccountRequest) (*emptypb.Empty, error) {
	if err := s.accountsHandler.DeleteAccount(req); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *server) ChangeAccountsBalance(ctx context.Context, req *proto.ChangeAccountsBalanceRequest) (*emptypb.Empty, error) {
	if err := s.accountsHandler.ChangeAccountsBalance(req); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *server) ChangeAccountsName(ctx context.Context, req *proto.ChangeAccountsNameRequest) (*emptypb.Empty, error) {
	if err := s.accountsHandler.ChangeAccountsName(req); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":1232")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	proto.RegisterBankServer(s, &server{accountsHandler: *accounts.New()})
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
