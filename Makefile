VERSION = 0.0.1

GIT_HASH := $(shell git show --format='%h' --no-patch)

BUILD_FLAG = -trimpath
LDFLAGS = -s -w -X 'github.com/byplayer/ttick/internal/cmd/ttick.version=$(VERSION) $(GIT_HASH)'

PROGRAM_NAME := ttick
PROGRAM := $(PROGRAM_NAME)

SRC = cmd/ttick/ttick.go

$(PROGRAM): $(SRC)
	go build -ldflags="$(LDFLAGS)" $(BUILD_FLAG) $(SRC)

.PHONY: build
build: $(PROGRAM)

.PHONY: clean
clean:
	rm -f $(PROGRAM)
