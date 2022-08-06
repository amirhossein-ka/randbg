DAEMON_OUT="./daemon"
DAEMON_SOURCE="./cmd/daemon/"
CONTROLLER_OUT="./controller"
CONTROLLER_SOURCE="./cmd/controller/"
GOPATH_BIN=""
LDFLAGS="-s -w"

.PHONY: all
all: .build_cmd .build_damon

.build_damon:
	go build -trimpath -ldflags=$(LDFLAGS) -o $(DAEMON_OUT) $(DAEMON_SOURCE)


.build_cmd:
	go build -trimpath -ldflags=$(LDFLAGS) -o $(CONTROLLER_OUT) $(CONTROLLER_SOURCE)



test:
	go test ./...

install: .build_damon .build_cmd
ifeq "$(origin GOPATH)" "undefined"
	@echo "GOPATH is not set. exiting..."
	exit 1
else
	mv $(CONTROLLER_OUT) "$(GOPATH)/bin"
	mv $(DAEMON_OUT) "$(GOPATH)/bin"
endif



clean:
	rm -f $(CONTROLLER_OUT)
	rm -f $(DAEMON_OUT)

