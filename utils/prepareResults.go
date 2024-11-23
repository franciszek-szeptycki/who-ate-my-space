package utils

import (
	"sort"
)

type DirInfo struct {
	AbsolutePath string
	Size         int64
}

type ExtendedDirInfo struct {
	AbsolutePath string
	SizeGB       float64
	SizeMB       float64
	SizeKB       float64
	SizeB        int64
}

func PrepareResults(dirSizeMap map[string]int64, limit int) []ExtendedDirInfo {
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

	var limitedDirInfos = make([]DirInfo, limit)
	copy(limitedDirInfos, dirInfos[:limit])

	return mapToExtendedDirInfo(limitedDirInfos)
}

func mapToExtendedDirInfo(dirInfos []DirInfo) []ExtendedDirInfo {
	extendedDirInfos := make([]ExtendedDirInfo, len(dirInfos))
	for id, dirInfo := range dirInfos {
		absPath, size := dirInfo.AbsolutePath, dirInfo.Size
		extendedDirInfo := ExtendedDirInfo{
			AbsolutePath: absPath,
			SizeGB:       roundToTwoDecimalPlaces(float64(size) / 1024 / 1024 / 1024),
			SizeMB:       roundToTwoDecimalPlaces(float64(size) / 1024 / 1024),
			SizeKB:       roundToTwoDecimalPlaces(float64(size) / 1024),
			SizeB:        size,
		}
		extendedDirInfos[id] = extendedDirInfo
	}

	return extendedDirInfos
}

func roundToTwoDecimalPlaces(num float64) float64 {
	return float64(int(num*100)) / 100
}
