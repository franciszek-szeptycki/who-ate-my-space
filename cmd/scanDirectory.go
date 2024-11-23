package cmd

import (
	"os"
	"path/filepath"
	"who-ate-my-space/cli"
)

func ScanDir(args cli.Args, fileMap *map[string]int64) {
	files, err := os.ReadDir(args.Path)
	if err != nil {
		return
	}

	for _, file := range files {
		absolutePath := filepath.Join(args.Path, file.Name())

		if file.IsDir() {
			ScanDir(cli.Args{
				Path:       absolutePath,
				IgnoreSize: args.IgnoreSize,
			}, fileMap)
		} else {
			fileInfo, err := os.Stat(absolutePath)
			if err != nil {
				continue
			}

			if fileInfo.Size() > args.IgnoreSize {
				(*fileMap)[absolutePath] = fileInfo.Size()
			}
		}
	}
}
