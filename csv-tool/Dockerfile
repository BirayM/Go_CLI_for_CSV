# --- Stage 1 : Compilation ---
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Copier les fichiers de dépendances et télécharger
COPY go.mod ./
# Si vous avez un go.sum, décommentez la ligne suivante :
# COPY go.sum ./
RUN go mod download

# Copier le reste du code source et compiler le binaire
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /csv-tool main.go

# --- Stage 2 : Image finale d'exécution ---
FROM alpine:3.19

WORKDIR /root/

# Copier le binaire compilé depuis l'étape précédente
COPY --from=builder /csv-tool .
# Copier le fichier CSV si votre binaire en a besoin localement
COPY --from=builder /app/data.csv .

# Exposer le port par défaut attendu par Cloud Run
ENV PORT=8080
EXPOSE 8080

CMD ["./csv-tool"]
