package main

import (
	"context"
	"fmt"
	"path/filepath"
	"time"

	"github.com/google/uuid"
	"github.com/mortalglitch/hashwood/internal/database"
)

func (cfg *appConfig) CommandIgnore(words []string) error {
	if len(words) > 1 {
		action := words[1]
		var directory string
		var targetFile string
		if len(words) > 2 {
			directory = filepath.Dir(words[2])
			targetFile = filepath.Base(words[2])
			if action == "add" {
				err := addFileToIgnore(targetFile, directory, cfg)
				if err != nil {
					return err
				}
			} else if action == "remove" {
				err := removeItemFromIgnore(targetFile, directory, cfg)
				if err != nil {
					return err
				}
			}
		}

		if action == "list" {
			err := ListIgnored(cfg)
			if err != nil {
				return err
			}
		}

	} // else show usage

	return nil
}

func addFileToIgnore(target string, directory string, cfg *appConfig) error {
	_, check, _ := CheckIfIgnored(target, directory, cfg)

	if check == false {
		ignoreEntry, err := cfg.db.CreateIgnoreEntry(context.Background(), database.CreateIgnoreEntryParams{
			ID:        uuid.New(),
			FileName:  target,
			Directory: directory,
			DateAdded: time.Now().UTC(),
		})
		if err != nil {
			return err
		}

		fmt.Printf("%s added to ignore list with id %s\n", ignoreEntry.FileName, ignoreEntry.ID)
	}
	return nil
}

func CheckIfIgnored(target string, directory string, cfg *appConfig) (uuid.UUID, bool, error) {
	ignoreEntry, _ := cfg.db.GetIgnoredItemByNameDirectory(context.Background(), database.GetIgnoredItemByNameDirectoryParams{
		FileName:  target,
		Directory: directory,
	})

	if ignoreEntry == (database.Ignorelist{}) {
		return uuid.UUID{}, false, nil
	}
	fmt.Printf("%s is on the ignore list.\n", ignoreEntry.FileName)
	return ignoreEntry.ID, true, nil
}

func ListIgnored(cfg *appConfig) error {
	ignoreList, err := cfg.db.GetIgnoreListByDateAdded(context.Background())
	if err != nil {
		return err
	}
	fmt.Printf("ID | Filename | Directory | Date Added\n")
	for _, item := range ignoreList {
		fmt.Printf("%s - %s - %s - %s\n", item.ID, item.FileName, item.Directory, item.DateAdded)
	}
	return nil
}

func removeItemFromIgnore(target string, directory string, cfg *appConfig) error {
	currentUUID, check, err := CheckIfIgnored(target, directory, cfg)
	if err != nil {
		return err
	}

	if check == true {
		deleteErr := cfg.db.DeleteIgnoreItemByID(context.Background(), currentUUID)
		if deleteErr != nil {
			return deleteErr
		}
	} else {
		fmt.Println("File doesn't exist in ignore list")
	}

	fmt.Printf("%s has been removed from the ignore list.", target)

	return nil
}
