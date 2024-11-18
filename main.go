package main

import (
	"fmt"
	"grodt/bitpanda"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	resp, err := bitpanda.GetTransactions("1ee32929-eee6-6184-9801-a29a90058262")
	if err != nil {
		panic(err)
	}
	fmt.Println(len(resp.Data))
	PrettyPrint(resp.Meta)
}
