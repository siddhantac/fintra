package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	ms := NewMemStore()

	tx, err := NewTransaction(23, time.Now(), true, string(TrCategoryEntertainment), string(TrTypeExpense), "movies", "Citibank")
	if err != nil {
		log.Fatal(err)
	}
	ms.Insert(tx)
	time.Sleep(time.Millisecond * 50)

	tx2, err := NewTransaction(11, time.Now(), true, string(TrCategoryMeals), string(TrTypeExpense), "foodpanda", "Citibank")
	if err != nil {
		log.Fatal(err)
	}
	ms.Insert(tx2)
	time.Sleep(time.Millisecond * 50)

	tx3, err := NewTransaction(12, time.Now(), true, string(TrCategoryMeals), string(TrTypeExpense), "deliveroo", "Citibank")
	if err != nil {
		log.Fatal(err)
	}
	ms.Insert(tx3)

	alltx := ms.GetAll()

	fmt.Println("no. of txn: ", len(alltx))
	for _, tx := range alltx {
		fmt.Printf("%+v\n", tx)
	}
}
