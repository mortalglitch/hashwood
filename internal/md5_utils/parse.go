package md5utils

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/mortalglitch/hashwood/internal/database"
	"github.com/mortalglitch/hashwood/internal/helpers"
)

type HashData struct {
	Filename string
	Hash     []byte
}

func ParseDirectory(target string, db *database.Queries) ([]HashData, error) {
	files, err := os.ReadDir(target)
	if err != nil {
		return nil, err
	}

	var results []HashData

	for _, file := range files {
		targetFile := filepath.Base(file.Name())
		_, check, _ := helpers.CheckIfIgnored(targetFile, target, db)

		info, err := os.Stat(target + "/" + file.Name())
		if err != nil {
			return nil, err
		}

		if info.IsDir() {
			fmt.Printf("Found directory %s skipping", file.Name())
		} else {
			if check == false {
				f, err := os.Open(target + "/" + file.Name())
				if err != nil {
					return nil, err
				}
				defer f.Close()

				h := md5.New()
				if _, err := io.Copy(h, f); err != nil {
					return nil, err
				}

				results = append(results, HashData{
					Filename: targetFile,
					Hash:     h.Sum(nil),
				})
			}
		}
	}

	return results, nil
}

func ParseFile(target string, db *database.Queries) ([]HashData, error) {
	file, err := os.Open(target)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var results []HashData

	targetFile := filepath.Base(file.Name())
	_, check, _ := helpers.CheckIfIgnored(targetFile, target, db)

	if check == false {
		h := md5.New()
		if _, err := io.Copy(h, file); err != nil {
			return nil, err
		}

		results = append(results, HashData{
			Filename: targetFile,
			Hash:     h.Sum(nil),
		})
	}

	return results, nil
}
