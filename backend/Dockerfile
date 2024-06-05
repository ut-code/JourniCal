FROM golang:1.22.4-alpine3.20
WORKDIR /usr/docker/JourniCal/backend/
COPY . .
RUN go mod tidy
CMD go run .
