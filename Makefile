BIN=rubiksImg
BINDIR=bin

build:
	go build -o $(BINDIR)/$(BIN) main.go

run:
	go run main.go

clean:
	rm $(BINDIR)/$(BIN)
