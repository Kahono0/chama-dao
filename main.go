package main

import (
	"fmt"
	"github.com/Kahono0/chama-dao/utils"
	"github.com/Kahono0/chama-dao/routes"
	"net/http"
)

func main() {
	utils.ConnectDB()
	fmt.Println("Connected to DB")
	mux := routes.SetupRoutes()
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	fmt.Println("Server is running on http://localhost:8080")
	server.ListenAndServe()

}
