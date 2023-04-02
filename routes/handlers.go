package routes

import (
	"fmt"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
	fmt.Println("Hello World")
}
