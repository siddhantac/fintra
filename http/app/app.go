package app

import (
	"log"
	"net/http"
)

func ShowHomePage(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte(pageHTML))
	if err != nil {
		log.Printf("error: %v\n", err)
	}
}
