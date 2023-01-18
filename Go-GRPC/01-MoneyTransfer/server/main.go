package main

import (
	pT "MoneyTransfer/transactions"
	"context"
	"errors"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	pT.UnimplementedTransactionsServer
}

func main() {
	s := grpc.NewServer()
	lis, err := net.Listen("tcp", ":8001")
	defer lis.Close()
	if err != nil {
		log.Fatalf("Port Not Available", err)
	}

	pT.RegisterTransactionsServer(s, &server{})
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to listen...\n", err)
	}
	defer s.GracefulStop()
}

func (s *server) MakeTransaction(ctx context.Context, in *pT.TransactionRequest) (*pT.TransactionResponse, error) {
	log.Println("Transaction Initiated...", ctx)
	log.Println("From : ", in.GetFrom())
	log.Println("From : ", in.GetTo())
	log.Println("From : ", in.GetAmount())
	log.Println("======================================================")
	return &pT.TransactionResponse{TransactionCompleted: true}, nil
}

func (s *server) CheckBalance(ctx context.Context, in *pT.EnquiryRequest) (*pT.EnquiryResponse, error) {
	log.Println("Checking Balance....")
	if in.GetPasscode() == "0102" {
		return &pT.EnquiryResponse{
			Balance: 106000.023,
		}, nil
	} else {
		return nil, errors.New("Incorrect Passcode.")
	}
}
