package main

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/mortalglitch/hashwood/internal/database"
)

func (cfg *appConfig) CommandHistory(words []string) error {
	if len(words) > 1 {
		resolveAbsolute, err := filepath.Abs(words[1])
		if err != nil {
			return err
		}
		directory := filepath.Dir(resolveAbsolute)
		targetFile := filepath.Base(resolveAbsolute)
		// Look up file from files DB to grab the ID then call history
		dbEntry := checkIfExist(targetFile, directory, cfg)
		if dbEntry == (database.File{}) {
			return fmt.Errorf("Unable to find database entry")
		}
		history, err := cfg.db.GetHistoryByFileID(context.Background(), dbEntry.ID)
		if err != nil {
			return err
		}
		fmt.Printf("Known history of %s:\n", targetFile)
		for _, item := range history {
			fmt.Printf("Previous Hash: %s Current Hash: %s Changed on %s\n", item.PreviousHash, item.CurrentHash, item.DateChange.String())
		}
		return nil
	}

	history, err := cfg.db.GetHistoryByDateChanged(context.Background())
	if err != nil {
		return err
	}

	for _, item := range history {
		fileName, err := cfg.db.GetFileNameByID(context.Background(), item.FileID)
		if err != nil {
			return err
		}
		fmt.Printf("File: %s Previous Hash: %s Current Hash: %s Changed on %s\n", fileName, item.PreviousHash, item.CurrentHash, item.DateChange.String())
	}

	return nil
}
