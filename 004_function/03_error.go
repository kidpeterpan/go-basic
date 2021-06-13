package main

import (
	"errors"
	"fmt"
	"log"
	"time"
)

const isDBReady = false // change this to false to see error
const balance = 2000
const dbRetryTimes = 2
const withdrawAmount = 2000 // withdrawAmount > balance also error

func getBalance() (int, error) {
	if !isDBReady {
		err := waitForDb()
		if err != nil {
			return 0, fmt.Errorf("getBalance: database is down : %v", err)
		}
	}
	return balance, nil
}

func waitForDb() error {
	for i := 0; i < 3; i++ {
		err := connectToDB(i)
		if err != nil {
			return fmt.Errorf("waitForDB: %v", err)
		}
		time.Sleep(time.Second)
	}
	return nil
}

func connectToDB(time int) error {
	if time == dbRetryTimes {
		return errors.New("connectToDB: maximum retry exceed")
	}
	log.Println("connecting to DB...")
	return nil
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
	money, err := withdraw(withdrawAmount)
	if err != nil {
		log.Fatal("main: ", err)
		return
	}
	fmt.Println("I got my money from bank:", money)
}

// note: เรื่องของ error จะต้องเอาไปจับกับ concept ต่างๆ เพิ่มเติมเช่น
// - stop program gracefully -> ถ้า error เราควร clear resource ให้อยู่ stat ที่ไม่สร้างความเสียหาย
// - reduce functionality -> fallback ไปใช้อะไรสักอย่างแทนได้ไหม?
// - ignore error -> บางอย่างไม่ได้สำคัญ อาจจะไม่ต้อง return error แค่ log ก็พอ
