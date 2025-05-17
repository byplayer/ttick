VERSION = $(shell cat VERSION | head -1)

GIT_HASH := $(shell git show --format='%h' --no-patch)

DEPENDS := VERSION

###
# ソースコードディレクトリ指定
###
SOURCEDIR     = ./cmd ./internal
TSOURCEDIR    = ./test

BUILD_FLAG = -trimpath
LDFLAGS = -s -w -X 'github.com/byplayer/ttick/internal/cmd/ttick.version=$(VERSION) $(GIT_HASH)'

PROGRAM_NAME := ttick
PROGRAM := $(PROGRAM_NAME)

###
# 処理部
###
# 1. サブディレクトリを含むディレクトリリストの生成
SRCDIRLIST  := $(shell find $(SOURCEDIR) -type d)
TSRCDIRLIST := $(shell find $(TSOURCEDIR) -type d)

# 2. 全てのgoファイルのリストの生成
SRCLIST     = $(foreach srcdir, $(SRCDIRLIST), $(wildcard $(srcdir)/*.go))
TSRCLIST    = $(foreach testsrcdir, $(TSRCDIRLIST), $(wildcard $(testsrcdir)/*.go))

$(PROGRAM): $(SRCLIST) $(DEPENDS)
	go build -ldflags="$(LDFLAGS)" $(BUILD_FLAG) $(SRC)

.PHONY: build
build: $(PROGRAM)

.PHONY: clean
clean:
	rm -f $(PROGRAM)
