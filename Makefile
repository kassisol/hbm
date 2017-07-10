TARGETS := $(shell ls scripts | grep -vE 'clean|dev|help|release')

TMUX := $(shell command -v tmux 2> /dev/null)

.dapper:
	@echo Downloading dapper
	@curl -sL https://releases.rancher.com/dapper/latest/dapper-`uname -s`-`uname -m|sed 's/v7l//'` > .dapper.tmp
	@@chmod +x .dapper.tmp
	@./.dapper.tmp -v
	@mv .dapper.tmp .dapper

.github-release:
	@echo Downloading github-release
	@curl -sL https://github.com/aktau/github-release/releases/download/v0.6.2/linux-amd64-github-release.tar.bz2 | tar xjO > .github-release.tmp
	@@chmod +x .github-release.tmp
	@./.github-release.tmp -v
	@mv .github-release.tmp .github-release

.tmass:
	@echo Downloading tmass
	@curl -sL https://github.com/juliengk/tmass/releases/download/0.3.0/tmass -o .tmass.tmp
	@@chmod +x .tmass.tmp
	@./.tmass.tmp version
	@mv .tmass.tmp .tmass

$(TARGETS): .dapper
	./.dapper $@

clean:
	@./scripts/clean

dev: .dapper .tmass
ifndef TMUX
	$(error "tmux is not available, please install it")
endif

	./.tmass load -l scripts/dev/tmux/ hbm
	tmux a -d -t hbm

help:
	@./scripts/help

release: .github-release
	./scripts/release

.DEFAULT_GOAL := ci

.PHONY: .dapper .github-release .tmass $(TARGETS) clean dev help release
