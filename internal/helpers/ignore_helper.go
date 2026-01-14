package helpers

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/mortalglitch/hashwood/internal/database"
)

func CheckIfIgnored(target string, directory string, db *database.Queries) (uuid.UUID, bool, error) {
	ignoreEntry, _ := db.GetIgnoredItemByNameDirectory(context.Background(), database.GetIgnoredItemByNameDirectoryParams{
		FileName:  target,
		Directory: directory,
	})

	if ignoreEntry == (database.Ignorelist{}) {
		return uuid.UUID{}, false, nil
	}
	fmt.Printf("%s is on the ignore list.\n", ignoreEntry.FileName)
	return ignoreEntry.ID, true, nil
}
