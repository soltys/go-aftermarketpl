package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/soltys/go-aftermarketpl"
)

func main() {
	key, _ := getenv("AFTERMARKETPL_KEY")
	secret, _ := getenv("AFTERMARKETPL_SECRET")
	client := aftermarketpl.New(key, secret)
	domainInfo, _ := client.DomainGet("soltysiak.it")

	fmt.Println(domainInfo.Name)
}

func getenv(key string) (string, error) {
	if value, ok := os.LookupEnv(key); ok {
		return value, nil
	}
	return "", errors.New("key was not found")
}
