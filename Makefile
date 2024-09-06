run:
	sqlc generate
	swag init
	mkdir -p ./bin
	go build -o ./bin/hoolo-bridge && ./bin/hoolo-bridge

debug:
	sqlc generate
	swag init
	air

deploy:
	sqlc generate
	swag init
	mkdir -p ./bin
	go run main.go -deploy