run:
    go run ./app/main.go

test:
    go test ./test/...

format:
    go fmt ./...

lint:
    golangci-lint run ./...

watch target:
    watchexec -e go,mod -c -r "just {{target}}"