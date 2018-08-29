TARGETS := $(shell ls scripts | grep -vE 'clean|dev|help|release|run-test')

TMUX := $(shell command -v tmux 2> /dev/null)

.PHONY: .dapper
.dapper:
	@echo Downloading dapper
	@curl -sL https://releases.rancher.com/dapper/latest/dapper-`uname -s`-`uname -m|sed 's/v7l//'` > .dapper.tmp
	@@chmod +x .dapper.tmp
	@./.dapper.tmp -v
	@mv .dapper.tmp .dapper

.PHONY: .github-release
.github-release:
	@echo Downloading github-release
	@curl -sL https://github.com/aktau/github-release/releases/download/v0.6.2/linux-amd64-github-release.tar.bz2 | tar xjO > .github-release.tmp
	@@chmod +x .github-release.tmp
	@./.github-release.tmp -v
	@mv .github-release.tmp .github-release

.PHONY: .tmass
.tmass:
	@echo Downloading tmass
	@curl -sL https://github.com/juliengk/tmass/releases/download/0.3.0/tmass -o .tmass.tmp
	@@chmod +x .tmass.tmp
	@./.tmass.tmp version
	@mv .tmass.tmp .tmass

.PHONY: $(TARGETS)
$(TARGETS): .dapper
	./.dapper $@

.PHONY: clean
clean:
	@./scripts/clean

.PHONY: dev
dev: .dapper .tmass
ifndef TMUX
	$(error "tmux is not available, please install it")
endif

	./.tmass load -l scripts/dev/tmux/ hbm
	tmux a -d -t hbm

.PHONY: help
help:
	@./scripts/help

.PHONY: release
release: .github-release
	./scripts/release

.PHONY: run-test
run-test:
	./scripts/run-test

.DEFAULT_GOAL := ci
