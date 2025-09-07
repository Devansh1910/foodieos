package getOutletFood

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

func StartServer() {
	// initialize DB (connect + migrate)
	InitDB()

	r := mux.NewRouter()
	r.HandleFunc("/api/getOutletFood", getOutletFoodHandler).Methods("POST")
	r.HandleFunc("/api/updateOutletFood", updateOutletFoodHandler).Methods("POST")

	// serve admin UI (index.html) and assets from ./web
	fs := http.FileServer(http.Dir("./web"))
	r.PathPrefix("/").Handler(fs)

	// Cloud Run provides PORT via env var
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // fallback for local
	}
	
	srv := &http.Server{
		Handler:      r,
		Addr:         ":" + port,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}
	

	log.Printf("listening on :%s\n", port)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
