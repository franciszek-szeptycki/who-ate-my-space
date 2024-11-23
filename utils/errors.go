package utils

import (
	"fmt"
	"os"
)

// HandleError - Obsługa błędów (logowanie błędów)
func HandleError(err error) {
	if os.IsPermission(err) {
		fmt.Printf("Brak uprawnień do pliku/katalogu: %v\n", err)
		return
	}
	fmt.Printf("Błąd: %v\n", err)
}
