package main

import (
	pT "MoneyTransfer/transactions"
	"context"
	"google.golang.org/grpc"
	"log"
)

const address = "localhost:8001"

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalln("Unable to Connect to Server...", err)
	}

	c := pT.NewTransactionsClient(conn)

	transactionResponse, err := c.MakeTransaction(context.Background(), &pT.TransactionRequest{
		From:   "Rahil",
		To:     "Chini",
		Amount: 500.11,
	})
	if err != nil {
		log.Fatalln("Transaction Failed...", err)
	}
	var status string
	if transactionResponse.GetTransactionCompleted() {
		status = "Successful"
	} else {
		status = "Unsuccessful"
	}
	log.Println("Transaction Status : ", status)

	log.Println("Checking Balance...")
	enquiryResponse, err := c.CheckBalance(context.Background(), &pT.EnquiryRequest{
		From:     "Rahil Jaiswal",
		Passcode: "0102",
	})
	if err != nil {
		log.Fatalln("Balance Check failed :", err)
	}
	log.Println("Current Balance :", enquiryResponse.GetBalance())
}
