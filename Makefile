MODULE_NAME = $(shell cat go.mod | grep "^module" | sed -e "s/module //g")

update_go_module:
	go mod tidy

install_toolkit: update_go_module
	@go install github.com/machinefi/w3bstream/pkg/depends/gen/cmd/...@latest

format: install_toolkit
	@toolkit fmt

build_examples:
	@cd examples && mkdir -pv wasms && rm -rf wasms/* && \
	for prj in * ; \
	do \
		if [ -d $$prj ] && [ $$prj != "wasms" ]; then \
			cd $$prj && echo "\033[32mbuilding $$prj ... \033[0m" ; \
			tinygo build -o $$prj.wasm --no-debug -target=wasi ; \
			mv $$prj.wasm ../wasms ; cd ..; \
			echo "\033[31mdone!\033[0m"; \
		fi \
	done
