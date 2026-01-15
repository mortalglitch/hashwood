package main

import (
	"context"
	"fmt"
	"path/filepath"
	"time"

	"github.com/google/uuid"
	"github.com/mortalglitch/hashwood/internal/database"
	md5utils "github.com/mortalglitch/hashwood/internal/md5_utils"
)

func (cfg *appConfig) CommandScan(words []string) error {
	if len(words) > 2 {
		scanType := words[1]
		if scanType == "file" {
			err := scanFile(words, cfg)
			if err != nil {
				return err
			}
		} else if scanType == "directory" {
			err := scanDirectory(words, cfg)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func checkIfExist(filename string, directory string, cfg *appConfig) database.File {
	exist, _ := cfg.db.GetFileByName(context.Background(), database.GetFileByNameParams{
		FileName:  filename,
		Directory: directory,
	})

	if exist == (database.File{}) {
		return exist
	} else {
		err := cfg.db.UpdateFileChecked(context.Background(), database.UpdateFileCheckedParams{
			UpdatedAt: time.Now().UTC(),
			ID:        exist.ID,
		})
		if err != nil {
			return exist
		}
	}
	return exist
}

func scanFile(words []string, cfg *appConfig) error {
	if len(words) > 2 {
		targetFile, err := filepath.Abs(words[2])
		if err != nil {
			return err
		}
		directory := filepath.Dir(targetFile)

		hashResults, err := md5utils.ParseFile(targetFile, cfg.db)
		if err != nil {
			return err
		}

		logErr := logHashResults(hashResults, directory, cfg)
		if logErr != nil {
			return logErr
		}
	}

	return nil
}

func scanDirectory(words []string, cfg *appConfig) error {
	directory, err := filepath.Abs(words[2])
	if err != nil {
		return err
	}

	hashResults, err := md5utils.ParseDirectory(directory, cfg.db)
	if err != nil {
		return err
	}

	logErr := logHashResults(hashResults, directory, cfg)
	if logErr != nil {
		return logErr
	}

	return nil
}

func logHashResults(hashResults []md5utils.HashData, directory string, cfg *appConfig) error {
	for _, hash := range hashResults {
		currentDBEntry := checkIfExist(hash.Filename, directory, cfg)

		if currentDBEntry == (database.File{}) {
			newEntry, err := cfg.db.CreateFileHash(context.Background(), database.CreateFileHashParams{
				ID:         uuid.New(),
				FileName:   hash.Filename,
				Directory:  directory,
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

			fmt.Printf("%s - %x added to database: %s\nHistory entry %s\n", hash.Filename, hash.Hash, newEntry.ID, history.ID)
		} else {
			if currentDBEntry.Hash != fmt.Sprintf("%x", hash.Hash) {
				fmt.Printf("Conflict Detected with %s\n", hash.Filename)
				err := cfg.ConflictDetected(currentDBEntry, hash)
				return err
			} else {
				fmt.Printf("File %s already exist and has not been added to the DB\n", hash.Filename)
			}
		}
	}

	return nil
}
