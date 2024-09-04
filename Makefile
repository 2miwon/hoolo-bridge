run:
	mkdir -p ./bin
	go build -o ./bin/hoolo-bridge && ./bin/hoolo-bridge

debug:
	air
