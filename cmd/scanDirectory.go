package cmd

import (
	"os"
	"path/filepath"
	"who-ate-my-space/utils"
)

func ScanDir(path string, dirSizeMap *map[string]int64) {
	processDirectory(path, dirSizeMap)
}

func processDirectory(path string, dirSizeMap *map[string]int64) {
	if utils.CheckIsPathExcluded(path) {
		return
	}

	totalSize := int64(0)
	dirs := listDirContents(path)

	for _, file := range dirs {
		absolutePath := filepath.Join(path, file.Name())

		if file.IsDir() {
			childDirSize := processSubdirectory(absolutePath, dirSizeMap)
			totalSize += childDirSize
		} else {
			fileSize := getFileSize(absolutePath)
			totalSize += fileSize
		}
	}

	(*dirSizeMap)[path] = totalSize
}

func listDirContents(path string) []os.DirEntry {
	dirs, err := os.ReadDir(path)
	if err != nil {
		return nil
	}
	return dirs
}

func processSubdirectory(path string, dirSizeMap *map[string]int64) int64 {
	processDirectory(path, dirSizeMap)
	return (*dirSizeMap)[path]
}

func getFileSize(path string) int64 {
	fileInfo, err := os.Stat(path)
	if err != nil || fileInfo.IsDir() {
		return 0
	}
	return fileInfo.Size()
}
