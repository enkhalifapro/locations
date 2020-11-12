# Go parameters
SERVER_BINARY=location
GEN_CODE=gen

all: build test
build: 
	go build -o $(SERVER_BINARY) main.go
test: 
	go test -v ./...
clean: 
	rm -f $(SERVER_BINARY)
	rm -rf $(GEN_CODE)
gen:
	goa gen locations/api/design
run:
	./$(SERVER_BINARY)
build-docker:
	docker build . -t locations
run-docker:
	docker run -p 9000:8000 locations