package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/mortalglitch/hashwood/internal/database"
	md5utils "github.com/mortalglitch/hashwood/internal/md5_utils"
)

func (cfg *appConfig) ConflictDetected(currentDBEntry database.File, newHash md5utils.HashData) error {
	// Log Conflict in history
	_, err := cfg.db.CreateHistoryEntry(context.Background(), database.CreateHistoryEntryParams{
		ID:           uuid.New(),
		PreviousHash: currentDBEntry.Hash,
		CurrentHash:  fmt.Sprintf("%x", newHash.Hash),
		DateChange:   time.Now().UTC(),
		FileID:       currentDBEntry.ID,
	})
	if err != nil {
		return err
	}

	// Update File database hash
	didUpdate := cfg.db.UpdateFileHash(context.Background(), database.UpdateFileHashParams{
		UpdatedAt: time.Now().UTC(),
		Hash:      fmt.Sprintf("%x", newHash.Hash),
		ID:        currentDBEntry.ID,
	})
	if didUpdate != nil {
		return didUpdate
	}

	// Report Conflict To Selected Channels
	// TODO

	return nil
}
