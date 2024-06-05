# How to run the server

1. install go (or use nix/docker to not pollute the environment)
2. cd into `JourniCal/backend`
3. run `go mod tidy`. this will install necessary packages and remove unnecessary ones automatically.
  alternatively, you can `go mod download`. this will download necessary packages without updating go.mod (I guess).
4. run `go run .`

you can also run `go build` to get an executable file.

# About Go

[Go](https://go.dev/) is a statically-typed compiled language intended to be as simple as possible.

To get started with Go, read https://go.dev/tour/welcome/1

For documentation, read https://go.dev/doc/
