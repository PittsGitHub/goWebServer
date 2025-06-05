package services

import (
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "The server is up")
}
