GOCDM=go
GOBUILD=$(GOCDM) build
GOMOD=$(GOCDM) mod
BINARY=discrete
GOHOME=~/go

all: clean tidy build 

build:
	env $(GOBUILD) -v -ldflags="-extldflags=-static" -o ${BINARY} cmd/discrete/discrete.go

tidy:
	$(GOMOD) tidy

move:
	mv ${BINARY} ${GOHOME}/bin/${BINARY}

clean:
	rm -f ${BINARY} ${GOHOME}/bin/${BINARY}
