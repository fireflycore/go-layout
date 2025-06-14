# 生成dto
.PHONY: dto
dto:
	goverter gen ./internal/data/dto

# 生成代码
.PHONY: generate
generate:
	buf generate
	make dto

# 初始化项目
.PHONY: init
init:
	make generate
	wire ./cmd/server
	go mod tidy

# 构建项目
.PHONY: build
build:
	make init
	go build -ldflags="-s -w" -o main ./cmd/server/

# 提交代码
.PHONY: push
push:
	git push gitee main
	git push github main

# 拉取代码
.PHONY: pull
pull:
	git pull --rebase gitee main