package main

import (
	"context"
	"fmt"
	"log"
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
			result, err := scanDirectory(words, cfg)
			if err != nil {
				return err
			}
			fmt.Println(result)
		}
	}

	return nil
}

func checkIfExist(filename string, directory string, cfg *appConfig) string {
	exist, _ := cfg.db.GetFileHashByName(context.Background(), database.GetFileHashByNameParams{
		FileName:  filename,
		Directory: directory,
	})

	if exist == (database.File{}) {
		return fmt.Sprint("Doesn't Exist")
	}
	return fmt.Sprint("Exist")
}

func scanFile() {

}

func scanDirectory(words []string, cfg *appConfig) (string, error) {
	var overallResult string

	// Add safety check
	targetDirectory := words[2]
	hashResults, err := md5utils.ParseDirectory(targetDirectory)
	if err != nil {
		log.Fatal(err)
	}

	for _, hash := range hashResults {
		fmt.Printf("%s - %x\n", hash.Filename, hash.Hash)
		exists := checkIfExist(hash.Filename, targetDirectory, cfg)
		if exists == "Doesn't Exist" {
			_, err := cfg.db.CreateFileHash(context.Background(), database.CreateFileHashParams{
				ID:         uuid.New(),
				FileName:   hash.Filename,
				Directory:  targetDirectory,
				CreatedAt:  time.Now().UTC(),
				UpdatedAt:  time.Now().UTC(),
				LastChange: time.Now().UTC(),
				Hash:       fmt.Sprintf("%x", hash.Hash),
			})
			if err != nil { // This can be fixed up
				overallResult = "Failed to process"
				return overallResult, err
			}
			overallResult = "Successfully added to DB"
		} else {
			fmt.Println("File already exist and has not been added to the DB")
			// Need to check if the file exist but the hash is different
		}
	}

	return overallResult, nil
}
