package cmd

import (
	"fmt"
	"runtime"
)

// DisplayResults - Wyświetla top N największych elementów
func DisplayResults(files []FileInfo, top int) {
	fmt.Printf("Największe pliki/katalogi (%s):\n", runtime.GOOS)
	for i := 0; i < len(files) && i < top; i++ {
		fmt.Printf("%d. %s - %.2f MB\n", i+1, files[i].Path, float64(files[i].Size)/1024/1024)
	}
}
