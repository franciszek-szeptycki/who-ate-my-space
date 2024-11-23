package cmd

import (
	"os"
	"path/filepath"
	"who-ate-my-space/utils"
)

// FileInfo - Struktura przechowująca ścieżkę i rozmiar
type FileInfo struct {
	Path string
	Size int64
}

// ScanDirectory - Rekurencyjnie skanuje katalog i zwraca listę plików/katalogów z rozmiarami
func ScanDirectory(root string) ([]FileInfo, error) {
	var results []FileInfo

	entries, err := os.ReadDir(root) // Pobierz zawartość katalogu
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		fullPath := filepath.Join(root, entry.Name())

		info, err := entry.Info()
		if err != nil {
			utils.HandleError(err) // Ignoruj błędy
			continue
		}

		if entry.IsDir() {
			// Rekurencyjne skanowanie podkatalogów
			subResults, err := ScanDirectory(fullPath)
			if err != nil {
				utils.HandleError(err)
				continue
			}
			results = append(results, subResults...)
		} else {
			// Dodaj plik
			results = append(results, FileInfo{
				Path: fullPath,
				Size: info.Size(),
			})
		}
	}

	return results, nil
}
