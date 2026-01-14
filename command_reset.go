package main

import (
	"context"
	"fmt"
)

func (cfg *appConfig) CommandReset() error {
	fmt.Println("Resetting Database")
	err := cfg.db.DeleteIgnoreList(context.Background())
	if err != nil {
		return err
	}

	histErr := cfg.db.DeleteHistory(context.Background())
	if histErr != nil {
		return histErr
	}

	filesErr := cfg.db.DeleteFiles(context.Background())
	if filesErr != nil {
		return filesErr
	}

	fmt.Println("Database has been reset.")

	return nil
}
