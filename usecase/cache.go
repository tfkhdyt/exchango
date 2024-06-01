package usecase

import (
	"fmt"
	"io/fs"
	"os"
	"path"
	"strings"
	"time"
)

func GetCachePath() (string, error) {
	userCacheDir, err := os.UserCacheDir()
	if err != nil {
		return "", fmt.Errorf("failed to read user cache dir: %w", err)
	}

	myCacheDir := path.Join(userCacheDir, "exchango")

	return myCacheDir, nil
}

func ClearCache(all, force bool) error {
	myCacheDir, err := GetCachePath()
	if err != nil {
		return err
	}

	// read all files in the cache directory
	entries, err := os.ReadDir(myCacheDir)
	if err != nil {
		return err
	}

	deletedFiles := []fs.DirEntry{}

	month := strings.ToLower(string(time.Now().Month().String()))
	time := time.Now().Format(time.DateOnly)

	for _, entry := range entries {
		if all {
			deletedFiles = append(deletedFiles, entry)
			continue
		}

		switch strings.HasPrefix(entry.Name(), "currencies-") {
		case true:
			if entry.Name() != fmt.Sprintf("currencies-%v.json", month) {
				deletedFiles = append(deletedFiles, entry)
			}
		case false:
			if !strings.Contains(entry.Name(), time) {
				deletedFiles = append(deletedFiles, entry)
			}
		}
	}

	var confirm string
	if !force {
		for _, entry := range deletedFiles {
			info, err := entry.Info()
			if err != nil {
				return fmt.Errorf(
					"failed to get info of %s: %w",
					entry.Name(),
					err,
				)
			}
			fmt.Printf("%v (%v KB)\n", entry.Name(), (info.Size() / 1024))
		}

		fmt.Print("\nClear cache? (y/n): ")

		if _, err = fmt.Scan(&confirm); err != nil {
			return fmt.Errorf("failed to read user input: %w", err)
		}
	}

	if strings.ToLower(confirm) == "y" || force {
		for _, entry := range deletedFiles {
			err := os.RemoveAll(path.Join(myCacheDir, entry.Name()))
			if err != nil {
				return fmt.Errorf(
					"failed to delete %s: %w",
					entry.Name(),
					err,
				)
			}
		}

		fmt.Println("Cache cleared successfully")
	}

	return nil
}
