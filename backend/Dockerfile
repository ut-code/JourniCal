FROM golang:1.22.4-alpine3.20
WORKDIR /usr/docker/JourniCal/backend/
COPY go.mod go.sum .
RUN --mount=type=cache,target=/go/pkg/mod/ go mod download
COPY . .
RUN --mount=type=cache,target=/go/pkg/mod/ go build .
CMD ./backend
