package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// 1. Récupération du port injecté par Cloud Run (défaut à 8080 en local)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// 2. Définition d'une route HTTP basique
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Service CSV Tool en ligne !")
	})

	// Optionnel : Route pour exécuter votre traitement CSV via HTTP
	http.HandleFunc("/run", func(w http.ResponseWriter, r *http.Request) {
		// Placez ici la logique de traitement de votre fichier CSV
		log.Println("Exécution du traitement CSV...")
		fmt.Fprintf(w, "Traitement CSV effectué avec succès.")
	})

	// 3. Démarrage du serveur HTTP sur ":PORT" (binding sur 0.0.0.0)
	addr := ":" + port
	log.Printf("Serveur en écoute sur %s...", addr)
	
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Échec du démarrage du serveur : %v", err)
	}
}
