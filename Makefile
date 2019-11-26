outfile = openc-ctl

.PHONY: build
build:
	@echo "Running build script"
	@sh ./scripts/build.sh
	@echo "Building go binary"
	@go build -o $(outfile) *.go
	@chmod +x $(outfile)
