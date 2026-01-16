package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func (cfg *appConfig) handlerReports(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	historyResults, err := cfg.db.GetHistoryByDateChanged(context.Background())
	if err != nil {
		return
	}

	w.Write([]byte(fmt.Sprint("<html><head><title>Hashwood Report Server</title><style>table, th, td { border:1px solid black;}</style></head><body><h1>Hashwood Report</h1><p>Current Information</p>")))
	w.Write([]byte(fmt.Sprint("<table><tr><th>File</th><th>Directory</th><th>Current Hash</th><th>Previous Hash</th><th>Date Changed</th></tr>")))
	for _, item := range historyResults {
		itemName, err := cfg.db.GetFileNameByID(context.Background(), item.FileID)
		if err != nil {
			return
		}

		itemDirectory, err := cfg.db.GetFileDirectoryByID(context.Background(), item.FileID)
		w.Write([]byte(fmt.Sprintf("<tr><td>%s</td><td>%s</td><td>%s</td><td>%s</td><td>%s</td></tr>", itemName, itemDirectory, item.CurrentHash, item.PreviousHash, item.DateChange.String())))
	}

	w.Write([]byte(fmt.Sprint("</table></body></html>")))
}

// Need to add a shutdown function.
func (cfg *appConfig) startServer(words []string, srv *http.Server) error {
	if len(words) > 1 {
		command := words[1]
		if command == "start" {
			fmt.Println("Server starting on http://localhost:8080/report")
			go func() {
				if err := srv.ListenAndServe(); err != http.ErrServerClosed {
					fmt.Printf("Server error: %v\n", err)
				}
			}()
		} else if command == "stop" {
			fmt.Println("Stopping server...")
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

			if err := srv.Shutdown(ctx); err != nil {
				fmt.Printf("Shutdown failed: %v\n", err)
			} else {
				fmt.Println("Server stopped gracefully.")
			}
			cancel()
		}
	}

	return nil
}
