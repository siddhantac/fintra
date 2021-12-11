package main

import (
	// "log"
	"net/http"
	// "time"

	"github.com/siddhantac/fintra/api"
	"github.com/siddhantac/fintra/service"

	// "github.com/siddhantac/fintra/domain"
	"github.com/siddhantac/fintra/infra/store"
	"github.com/siddhantac/fintra/repository"
)

// func main() {
// 	ms := store.NewMemStore()

// 	tx, err := domain.NewTransaction(23, time.Now(), true, string(domain.TrCategoryEntertainment), string(domain.TrTypeExpense), "movies", "Citibank")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	ms.Insert(tx.ID, tx)

// 	tx2, err := domain.NewTransaction(11, time.Now(), true, string(domain.TrCategoryMeals), string(domain.TrTypeExpense), "foodpanda", "Citibank")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	ms.Insert(tx2.ID, tx2)

// 	tx3, err := domain.NewTransaction(12, time.Now(), true, string(domain.TrCategoryMeals), string(domain.TrTypeExpense), "deliveroo", "Citibank")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	ms.Insert(tx3.ID, tx3)

// 	alltx := ms.GetAll()

// 	fmt.Println("no. of txn: ", len(alltx))
// 	for _, tx := range alltx {
// 		fmt.Printf("%v\n", tx)
// 	}
// }

func main() {
	// tx, err := domain.NewTransaction(23, time.Now(), true, string(domain.TrCategoryEntertainment), string(domain.TrTypeExpense), "movies", "Citibank")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	txnRepo := repository.NewTransactionRepository(store.NewMemStore())
	svc := service.NewService(txnRepo)
	h := api.NewHandler(svc)

	http.HandleFunc("/transaction", h.CreateTransaction)
	http.ListenAndServe("", nil)
}
