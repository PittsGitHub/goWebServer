package main

import (
	"fmt"
	"net/http"

	"github.com/PittsGitHub/goWebServer/services"
	"github.com/PittsGitHub/goWebServer/types"
)

func main() {

	config := services.LoadConfig(types.ConfigFileName)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		services.Handler(w, r)
	})

	url := fmt.Sprintf("http://localhost:%s", config.Port)
	fmt.Printf("server running on:")
	fmt.Printf("\n%s\n", url)
	http.ListenAndServe(":"+config.Port, nil)
}
