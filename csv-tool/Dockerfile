# Stage 1 : Compilation
FROM golang:1.25.11-alpine AS builder

WORKDIR /src

# Copie et téléchargement des dépendances
COPY go.mod ./
# COPY go.sum ./  # Décommentez si vous utilisez go.sum
RUN go mod download

# Copie du code source et compilation du binaire Linux
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/csv-tool main.go

# Stage 2 : Image d'exécution minimale
FROM alpine:3.19

WORKDIR /root/

# Copie du binaire depuis le stage builder
COPY --from=builder /app/csv-tool .
# Copie du fichier data.csv si le binaire en a besoin
COPY --from=builder /src/data.csv .

EXPOSE 8080

CMD ["./csv-tool"]
