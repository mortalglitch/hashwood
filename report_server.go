package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

func (cfg *appConfig) handlerReports(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	historyResults, err := cfg.db.GetHistoryByDateChanged(context.Background())
	if err != nil {
		return
	}

	w.Write([]byte(fmt.Sprint("<html><body><h1>Hashwood Report</h1><p>Current Information</p>")))

	for _, item := range historyResults {
		itemName, err := cfg.db.GetFileNameByID(context.Background(), item.FileID)
		if err != nil {
			return
		}

		itemDirectory, err := cfg.db.GetFileDirectoryByID(context.Background(), item.FileID)
		w.Write([]byte(fmt.Sprintf("<p>File: %s Directory: %s Current Hash: %s Previous Hash: %s Last Change Detected: %s", itemName, itemDirectory, item.CurrentHash, item.PreviousHash, item.DateChange.String())))
	}

	w.Write([]byte(fmt.Sprint("</body></html>")))
}

// Need to add a shutdown function.
func (cfg *appConfig) startServer() error {
	const filepathRoot = "."
	const port = "8080"
	mux := http.NewServeMux()

	mux.HandleFunc("GET /report", cfg.handlerReports)
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
	log.Fatal(srv.ListenAndServe())
	return nil
}
