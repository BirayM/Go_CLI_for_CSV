package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

// readCSV ouvre et lit un fichier CSV, puis renvoie son contenu sous forme [][]string.
func readCSV(filePath string) ([][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("impossible d'ouvrir le fichier : %w", err)
	}
	// On s'assure que le fichier sera fermé à la fin de la fonction
	defer file.Close()

	reader := csv.NewReader(file)
	
	// ReadAll lit l'intégralité du fichier CSV en une seule fois
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("erreur lors de la lecture du CSV : %w", err)
	}

	return records, nil
}

func main() {
	// Fichier de test temporaire
	filename := "data.csv"

	fmt.Println("Lecture du fichier...", filename)
	data, err := readCSV(filename)
	if err != nil {
		fmt.Printf("Erreur : %v\n", err)
		return
	}

	fmt.Printf("Succès ! %d lignes lues (en-tête compris).\n\n", len(data))

	// Affichage des premières lignes pour vérifier
	for i, row := range data {
		if i > 5 { // On limite l'affichage aux 5 premières lignes
			fmt.Println("...")
			break
		}
		fmt.Println(row)
	}
}
