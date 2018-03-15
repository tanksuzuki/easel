.PHONY: help
.DEFAULT_GOAL := help

release: ## release build
	mkdir -p out
	rm -rf out/*
	GOOS=linux GOARCH=amd64 go build -o easel
	zip out/easel_linux_amd64.zip easel
	rm -rf easel
	GOOS=linux GOARCH=386 go build -o easel
	zip out/easel_linux_386.zip easel
	rm -rf easel
	GOOS=linux GOARCH=arm go build -o easel
	zip out/easel_linux_arm.zip easel
	rm -rf easel
	GOOS=darwin GOARCH=amd64 go build -o easel
	zip out/easel_macosx_amd64.zip easel
	rm -rf easel
	GOOS=darwin GOARCH=386 go build -o easel
	zip out/easel_macosx_386.zip easel
	rm -rf easel
	GOOS=windows GOARCH=amd64 go build -o easel.exe
	zip out/easel_windows_amd64.zip easel.exe
	rm -rf easel.exe
	GOOS=windows GOARCH=386 go build -o easel.exe
	zip out/easel_windows_386.zip easel.exe
	rm -rf easel.exe

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
