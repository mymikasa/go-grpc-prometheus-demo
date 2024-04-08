start:
	GOOS=linux GOARCH=amd64 go build -o ./server/server ./server/server.go
	GOOS=linux GOARCH=amd64 go build -o ./client/client ./client/client.go
	docker compose up -d