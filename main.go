package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"

	"github.com/mortalglitch/hashwood/internal/database"
	inputoutput "github.com/mortalglitch/hashwood/internal/input_output"

	"github.com/joho/godotenv"
)

type appConfig struct {
	db *database.Queries
}

func main() {
	fmt.Println("Starting Hashwood")
	fmt.Println("Use help for information about available commands.")
	godotenv.Load()

	dbURL := os.Getenv("DB_URL")

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("error connecting to db: %v", err)
	}

	dbQueries := database.New(db)

	cfg := appConfig{
		db: dbQueries,
	}

	asm := NewAutoScanManager()

	for {
		result := inputoutput.GetInput()

		if len(result) == 0 {
			continue
		} else if result[0] == "scan" {
			err := cfg.CommandScan(result)
			if err != nil {
				log.Fatal(err)
			}
		} else if result[0] == "help" {
			inputoutput.PrintHelp()
		} else if result[0] == "autoscan" {
			asm.CommandAutoScan(result, &cfg)
		} else if result[0] == "history" {
			err := cfg.CommandHistory(result)
			if err != nil {
				log.Fatal(err)
			}
		} else if result[0] == "ignore" {
			err := cfg.CommandIgnore(result)
			if err != nil {
				log.Fatal(err)
			}
		} else if result[0] == "reset" {
			err := cfg.CommandReset()
			if err != nil {
				log.Fatal(err)
			}
		} else if result[0] == "server" {
			err := cfg.CommandServer(result)
			if err != nil {
				log.Fatal(err)
			}
		} else if result[0] == "quit" {
			if activeServer != nil {
				stopServer()
			}
			break
		}
	}
}
