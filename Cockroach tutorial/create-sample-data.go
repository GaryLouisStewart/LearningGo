package main

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Account struct {
	ID      int `gorm:"primary_key"`
	Balance int
}

func main() {

	const addr = "postgresql://dbadmin@localhost:26257/bank?sslmode=disable"
	db, err := gorm.Open("postgres", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.AutoMigrate(&Account{})

	db.Create(&Account{ID: 1, Balance: 1000})
	db.Create(&Account{ID: 2, Balance: 1500})
	db.Create(&Account{ID: 3, Balance: 2000})
	db.Create(&Account{ID: 4, Balance: 2500})
	db.Create(&Account{ID: 5, Balance: 4000})
	db.Create(&Account{ID: 6, Balance: 10000})

	var accounts []Account
	db.Find(&accounts)
	fmt.Println("Initial balances:")
	for _, account := range accounts {
		fmt.Printf("%d %d\n", account.ID, account.Balance)
	}

}
