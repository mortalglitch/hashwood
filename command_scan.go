package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/mortalglitch/hashwood/internal/database"
	md5utils "github.com/mortalglitch/hashwood/internal/md5_utils"
)

func (cfg *appConfig) CommandScan(words []string) error {
	if len(words) > 2 {
		scanType := words[1]
		if scanType == "file" {
			fmt.Println("Add Functionality")
		} else if scanType == "directory" {
			err := scanDirectory(words, cfg)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func checkIfExist(filename string, directory string, cfg *appConfig) (database.File, string) {
	exist, _ := cfg.db.GetFileByName(context.Background(), database.GetFileByNameParams{
		FileName:  filename,
		Directory: directory,
	})

	if exist == (database.File{}) {
		return exist, fmt.Sprint("Doesn't Exist")
	} else {
		err := cfg.db.UpdateFileChecked(context.Background(), database.UpdateFileCheckedParams{
			UpdatedAt: time.Now().UTC(),
			ID:        exist.ID,
		})
		if err != nil {
			return exist, "Error updating timestamp"
		}
	}
	return exist, fmt.Sprint("Exist")
}

func scanFile() {
	// TODO add single file functionality
}

func scanDirectory(words []string, cfg *appConfig) error {
	targetDirectory := words[2]
	hashResults, err := md5utils.ParseDirectory(targetDirectory)
	if err != nil {
		return err
	}

	for _, hash := range hashResults {
		currentDBEntry, exists := checkIfExist(hash.Filename, targetDirectory, cfg)
		if exists == "Doesn't Exist" {
			newEntry, err := cfg.db.CreateFileHash(context.Background(), database.CreateFileHashParams{
				ID:         uuid.New(),
				FileName:   hash.Filename,
				Directory:  targetDirectory,
				CreatedAt:  time.Now().UTC(),
				UpdatedAt:  time.Now().UTC(),
				LastChange: time.Now().UTC(),
				Hash:       fmt.Sprintf("%x", hash.Hash),
			})
			if err != nil {
				return err
			}

			history, err := cfg.db.CreateHistoryEntry(context.Background(), database.CreateHistoryEntryParams{
				ID:           uuid.New(),
				PreviousHash: "none",
				CurrentHash:  newEntry.Hash,
				DateChange:   time.Now().UTC(),
				FileID:       newEntry.ID,
			})
			if err != nil {
				return err
			}

			fmt.Printf("%s - %x added to database and history entry %s\n", hash.Filename, hash.Hash, history.ID)
		} else {
			fmt.Println("File already exist and has not been added to the DB")
			if currentDBEntry.Hash != fmt.Sprintf("%x", hash.Hash) {
				fmt.Println("Conflict Detected")
				err := cfg.ConflictDetected(currentDBEntry, hash)
				return err
			}
		}
	}

	return nil
}
