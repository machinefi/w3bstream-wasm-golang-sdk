MODULE_NAME = $(shell cat go.mod | grep "^module" | sed -e "s/module //g")

update_go_module:
	go mod tidy

install_toolkit: update_go_module
	@go install github.com/machinefi/w3bstream/pkg/depends/gen/cmd/...@latest

format: install_toolkit
	@toolkit fmt

build_wasms:
	@cd examples && make build