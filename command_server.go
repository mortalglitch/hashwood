package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/mortalglitch/hashwood/internal/helpers"
)

var activeServer *http.Server

func (cfg *appConfig) handlerReports(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	historyResults, err := cfg.db.GetHistoryByDateChanged(context.Background())
	if err != nil {
		return
	}

	w.Write([]byte(`<html><head><title>Hashwood Report Server</title><meta http-equiv="refresh" content="60"/><style>table, th, td { border:1px solid black;}</style></head>
		<body><h1>Hashwood Report</h1><p>Current Information(changes and todays events highlighted in yellow)</p>`))
	w.Write([]byte("<table><tr><th>File</th><th>Directory</th><th>Current Hash</th><th>Previous Hash</th><th>Date Changed</th></tr>"))
	for _, item := range historyResults {
		itemName, err := cfg.db.GetFileNameByID(context.Background(), item.FileID)
		if err != nil {
			return
		}

		itemDirectory, err := cfg.db.GetFileDirectoryByID(context.Background(), item.FileID)
		isToday := helpers.CheckIfToday(item.DateChange)

		var htmlCurrentHash string
		if item.PreviousHash != "none" {
			htmlCurrentHash = fmt.Sprintf(`<td style="background-color: yellow;">%s</td>`, item.CurrentHash)
		} else {
			htmlCurrentHash = fmt.Sprintf(`<td>%s</td>`, item.CurrentHash)
		}
		// Cleanup note: This could be cleaned up with a different type of builder vs doing it all at once.
		if isToday {
			w.Write([]byte(fmt.Sprintf(`<tr><td>%s</td><td>%s</td>%s<td>%s</td><td style="background-color: yellow;">%s</td></tr>`, itemName, itemDirectory, htmlCurrentHash, item.PreviousHash, item.DateChange.String())))
		} else {
			w.Write([]byte(fmt.Sprintf(`<tr><td>%s</td><td>%s</td>%s<td>%s</td><td>%s</td></tr>`, itemName, itemDirectory, htmlCurrentHash, item.PreviousHash, item.DateChange.String())))
		}
	}

	w.Write([]byte("</table></body></html>"))
}

func (cfg *appConfig) CommandServer(words []string) error {
	const filepathRoot = "."
	const port = "8080"

	if len(words) > 1 {
		command := words[1]
		if command == "start" {
			if activeServer != nil {
				fmt.Println("Error: Server is already running.")
				return nil
			}
			mux := http.NewServeMux()

			mux.HandleFunc("GET /report", cfg.handlerReports)
			activeServer = &http.Server{
				Addr:    ":" + port,
				Handler: mux,
			}

			fmt.Println("Server starting on http://localhost:8080/report")
			go func() {
				if err := activeServer.ListenAndServe(); err != http.ErrServerClosed {
					fmt.Printf("Server error: %v\n", err)
					activeServer = nil
				}
			}()
		} else if command == "stop" {
			stopServer()
		}
	}

	return nil
}

func stopServer() {
	if activeServer == nil {
		fmt.Println("Error: No server is currently running.")
		return
	}

	fmt.Println("Stopping server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := activeServer.Shutdown(ctx); err != nil {
		fmt.Printf("Shutdown failed: %v\n", err)
	} else {
		fmt.Println("Server stopped gracefully.")
	}

	activeServer = nil
}
