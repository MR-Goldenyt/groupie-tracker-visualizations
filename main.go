package main

import (
	"fmt"
	"log"
	"net/http"

	// "os"

	backend "groupie-tracker-visualizations/project/backEnd"
	// "groupie-tracker/project/backEnd/helpers"
)

func main() {
	// if errs := helpers.CheckRequiredAssets(); errs != nil {
	// 	for _, err := range errs {
	// 		log.Println("Server failed:", err)
	// 	}
	// 	os.Exit(1)
	// }

	addr := ":8080"

	// Serve static assets from template and static under /static/
	http.Handle("/static/", backend.AssetHandler())

	// Print and start server
	fmt.Printf("🚀 Server listening at http://localhost%s\n", addr)
	log.Printf("Starting server on %s", addr[1:])

	// Dynamic routes
	http.HandleFunc("/", backend.Handler)

	// Run server
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
