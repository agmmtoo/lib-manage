
.PHONY: confirm
confirm:
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]

## run/api: run the cmd/api application
.PHONY: run/dev
run/dev:
	air -c .air.toml dev

## audit: tidy dependencies and format, vet and test all code
.PHONY: audit
audit: 
	@echo 'Tidying and verifying module dependencies...'
	go mod tidy
	go mod verify
	@echo 'Formatting code...'
	go fmt ./...
	@echo 'Vetting code...'
	go vet ./...
	staticcheck ./...
	@echo 'Running tests...'
	go test -race -vet=off ./...

## build/api: build the cmd/api application
.PHONY: build/lib
build/lib:
	@echo 'Building cmd/lib...'
	go build -o ./lib ./cmd/libraryapp

.PHONY: kill
kill:
	sudo lsof -t -i :8080 | xargs kill
