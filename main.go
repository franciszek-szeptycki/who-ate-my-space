package main

import (
	"flag"
	"log"
	"who-ate-my-space/cmd"
)

func main() {
	// Flagi CLI
	path := flag.String("path", ".", "Ścieżka do katalogu do skanowania")
	top := flag.Int("top", 10, "Liczba największych elementów do wyświetlenia")
	flag.Parse()

	// Skanowanie katalogu
	results, err := cmd.ScanDirectory(*path)
	if err != nil {
		log.Fatalf("Błąd podczas skanowania: %v", err)
	}

	// Sortowanie wyników
	cmd.SortBySize(results)

	// Wyświetlanie wyników
	cmd.DisplayResults(results, *top)
}
