package utils

import (
	"sort"
)

type DirInfo struct {
	AbsolutePath string
	Size         int64
}

func PrepareResults(dirSizeMap map[string]int64, limit int) []DirInfo {
	dirInfos := make([]DirInfo, len(dirSizeMap))
	for absPath, size := range dirSizeMap {
		dirInfo := DirInfo{
			AbsolutePath: absPath,
			Size:         size,
		}
		dirInfos = append(dirInfos, dirInfo)
	}

	sort.Slice(dirInfos, func(i, j int) bool {
		return dirInfos[i].Size > dirInfos[j].Size
	})

	return dirInfos[:limit]
}
