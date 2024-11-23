package utils

import (
	"sort"
)

type FileInfo struct {
	AbsolutePath string
	Size         int64
}

func GetHeaviestFiles(dirSizeMap map[string]int64, limit int) []FileInfo {
	fileInfos := make([]FileInfo, len(dirSizeMap))
	for absPath, size := range dirSizeMap {
		fileInfo := FileInfo{
			AbsolutePath: absPath,
			Size:         size,
		}
		fileInfos = append(fileInfos, fileInfo)
	}

	sort.Slice(fileInfos, func(i, j int) bool {
		return fileInfos[i].Size > fileInfos[j].Size
	})

	return fileInfos[:limit]
}
