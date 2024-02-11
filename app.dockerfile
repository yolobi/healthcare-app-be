FROM postgres:latest

# Install postgis extension
RUN apt-get update \
    && apt-get install -y postgis \
    && rm -rf /var/lib/apt/lists/*

FROM golang:1.18.10-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o migrate ./cmd/healthcare/database/migrate/
RUN CGO_ENABLED=0 GOOS=linux go build -o seed ./cmd/healthcare/database/seed/
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/healthcare/cmd/
RUN CGO_ENABLED=0 GOOS=linux go build -o drop ./cmd/healthcare/database/drop/

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/migrate /app/migrate
COPY --from=builder /app/seed /app/seed
COPY --from=builder /app/main /app/main
COPY --from=builder /app/drop /app/drop
COPY --from=builder /app/files/yaml/configs/healthcare.docker.yml /app/files/yaml/configs/healthcare.docker.yml
COPY --from=builder /app/files/csv/. /app/files/csv/.
COPY --from=builder /app/files/env/. /app/files/env/.
COPY --from=builder /app/files/cloud/. /app/files/cloud/.

EXPOSE 8080

CMD ./migrate -config /app/files/yaml/configs/healthcare.docker.yml && ./main -config /app/files/yaml/configs/healthcare.docker.yml
