default: start

start: build
	make serve

build:
	go build -o ./dist/backend -ldflags="-s -w" -trimpath .

serve:
	./dist/backend

full-test:
	go test ./test/...

test: vet staticcheck test-pkg test-app dry-build mutable-test

vet:
	go vet ./...
staticcheck:
	go run honnef.co/go/tools/cmd/staticcheck@latest ./...
test-pkg:
	go test ./pkg/...
test-app:
	ENV_FILE=test go test ./app/...
dry-build:
	go build -n 2>/dev/null
mutable-test:
	scripts/mutable-check.sh

database:
	docker run \
	  -e POSTGRES_PASSWORD=password \
	  -e POSTGRES_USER=user \
	  -e POSTGRES_DB=database \
	  -p 5432:5432 \
	  postgres:alpine
