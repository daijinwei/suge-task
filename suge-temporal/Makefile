# Define project
BINARY_NAME := suge-task
BINARY_DIR := ./bin
CREATE_WORK:= create-work
PRODUCE_START_WORK:= produce-start-work
CONSUME_START_WORK:= consume-start-work
SCHEDULE_WORK:=schedule-work

GO := go

# Define flags
BUILD_FLAGS := -v

# Define OS arch
OS := $(if $(GOOS),$(GOOS),$(shell $(GO) env GOOS))
ARCH := $(if $(GOARCH),$(GOARCH),$(shell $(GO) env GOARCH))

all: create-work start-produce-work start-consume-work

# build work
create-work:
	$(GO) build $(BUILD_FLAGS) -o ./bin/$(CREATE_WORK)_$(OS)_$(ARCH)  ./worker/main.go

# build produce start worker
start-produce-work:
	$(GO) build $(BUILD_FLAGS) -o ./bin/$(PRODUCE_START_WORK)_$(OS)_$(ARCH) ./start/produce_main.go

# build consumer start worker
start-consume-work:
	$(GO) build $(BUILD_FLAGS) -o ./bin/$(CONSUME_START_WORK)_$(OS)_$(ARCH) ./start/consumer_main.go

# clean
clean:
	rm -rf $(BINARY_DIR)/*
