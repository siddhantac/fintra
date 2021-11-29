package main

import (
	"fmt"
	"log"
	"time"

	"github.com/siddhantac/fintra/domain"
	"github.com/siddhantac/fintra/infra/store"
)

func main() {
	ms := store.NewMemStore()

	tx, err := domain.NewTransaction(23, time.Now(), true, string(domain.TrCategoryEntertainment), string(domain.TrTypeExpense), "movies", "Citibank")
	if err != nil {
		log.Fatal(err)
	}
	ms.Insert(tx)

	tx2, err := domain.NewTransaction(11, time.Now(), true, string(domain.TrCategoryMeals), string(domain.TrTypeExpense), "foodpanda", "Citibank")
	if err != nil {
		log.Fatal(err)
	}
	ms.Insert(tx2)

	tx3, err := domain.NewTransaction(12, time.Now(), true, string(domain.TrCategoryMeals), string(domain.TrTypeExpense), "deliveroo", "Citibank")
	if err != nil {
		log.Fatal(err)
	}
	ms.Insert(tx3)

	alltx := ms.GetAll()

	fmt.Println("no. of txn: ", len(alltx))
	for _, tx := range alltx {
		fmt.Printf("%v\n", tx)
	}
}
