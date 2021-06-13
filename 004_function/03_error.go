package main

import (
	"errors"
	"fmt"
	"log"
)

const isDBReady = true
const balance = 2000

func getBalance() (int, error) {
	if !isDBReady {
		return 0, errors.New("getBalance: database is down")
	}
	return balance, nil
}

// withdraw is example error propagation func
func withdraw(amount int) (int, error) {
	balance, err := getBalance()
	if err != nil {
		return 0, fmt.Errorf("withdraw: %v", err)
	}
	if amount > balance {
		return 0, errors.New("withdraw: insufficient fund")
	}
	return amount, nil
}

func main() {
	// note: error ใน go เป็นแค่ค่าๆ นึง
	// scenario: actor -> bank application -> database
	money, err := withdraw(2001)
	if err != nil {
		log.Fatal("main: ", err)
		return
	}
	fmt.Println("I got my money from bank:", money)
}
