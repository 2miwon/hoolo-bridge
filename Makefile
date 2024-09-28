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
	swag init
	mkdir -p ./bin
	go build -o ./bin/hoolo-bridge
	nohup ./bin/hoolo-bridge -deploy > /dev/null 2>&1 &

setup: 
	sqlc generate
	swag init
	mkdir -p ./bin
	go build -o ./bin/hoolo-bridge && nohup ./bin/hoolo-bridge &