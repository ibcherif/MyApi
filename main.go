package main

import (
	"cherif.com/myApi/appDoc"
	"log"
	"net/http"
)

func main() {
	//ecouter sur le serveur sur le port 8000
	err := http.ListenAndServe(":8000", appDoc.Handlers())

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
