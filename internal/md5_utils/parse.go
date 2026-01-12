package md5utils

import (
	"crypto/md5"
	"io"
	"os"
)

type HashData struct {
	Filename string
	Hash     []byte
}

func ParseDirectory(target string) ([]HashData, error) {
	files, err := os.ReadDir(target)
	if err != nil {
		return nil, err
	}

	var results []HashData

	for _, file := range files {
		f, err := os.Open(target + file.Name())
		if err != nil {
			return nil, err
		}
		defer f.Close()

		h := md5.New()
		if _, err := io.Copy(h, f); err != nil {
			return nil, err
		}

		results = append(results, HashData{
			Filename: f.Name(),
			Hash:     h.Sum(nil),
		})
	}

	return results, nil
}

func ParseFile(target string) ([]HashData, error) {
	file, err := os.Open(target)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var results []HashData

	h := md5.New()
	if _, err := io.Copy(h, file); err != nil {
		return nil, err
	}

	results = append(results, HashData{
		Filename: file.Name(),
		Hash:     h.Sum(nil),
	})

	return results, nil
}
