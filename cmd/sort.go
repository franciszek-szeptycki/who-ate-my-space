package cmd

import (
	"sort"
)

// SortBySize - Sortuje listę wyników według rozmiaru malejąco
func SortBySize(files []FileInfo) {
	sort.Slice(files, func(i, j int) bool {
		return files[i].Size > files[j].Size
	})
}
