package core

import (
	"os"
	"strings"
)

func DirReader(path string, showhidden bool) ([]string, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	filesList := make([]string, 0, len(entries)+2)

	if showhidden{
		filesList = append(filesList, ".","..")
	}

	for _, entry := range entries {
		name := entry.Name()

		if !showhidden && strings.HasPrefix(name, ".") {
			continue
		}
		filesList = append(filesList, name)
	}
	return filesList, nil
}
