install_easyjson:
	@if [ ! -f "$$GOBIN/easyjson" ] ; \
	then \
	  	echo "install easyjson to $$GOBIN" ; \
		go get github.com/mailru/easyjson && go install github.com/mailru/easyjson/...@latest ; \
	fi;

gen_wasm_json: install_easyjson
	@cd easyjson/model && go generate ./...

easyjson_wasm: gen_wasm_json
	@cd easyjson && tinygo build -o easyjson.wasm -scheduler=none --no-debug -target=wasi parse_json.go
	@mkdir -pv wasms && mv easyjson/easyjson.wasm wasms
	@echo built to wasms/easyjson.wasm

log_wasm:
	@cd log && tinygo build -o log.wasm -scheduler=none --no-debug -target=wasi log.go
	@mkdir -pv wasms && mv log/log.wasm wasms
	@echo built to wasms/log.wasm

gjson_wasm:
	@cd gjson && tinygo build -o gjson.wasm -scheduler=none --no-debug -target=wasi parse_json.go
	@mkdir -pv wasms && mv gjson/gjson.wasm wasms
	@echo built to wasms/gjson.wasm

token_distribute_wasm:
	@cd token_distribute && tinygo build -o token_distribute.wasm -scheduler=none --no-debug -target=wasi token_distribute.go
	@mkdir -pv wasms && mv token_distribute/token_distribute.wasm wasms
	@echo built to wasms/token_distribute.wasm

word_count_wasm:
	@cd word_count && tinygo build -o word_count.wasm -scheduler=none --no-debug -target=wasi word_count.go
	@mkdir -pv wasms && mv word_count/word_count.wasm wasms
	@echo built to wasms/word_count.wasm

word_count_v2_wasm:
	@cd word_count_v2 && tinygo build -o word_count_v2.wasm -scheduler=none --no-debug -target=wasi word_count_v2.go
	@mkdir -pv wasms && mv word_count_v2/word_count_v2.wasm wasms
	@echo built to wasms/word_count_v2.wasm

build: easyjson_wasm log_wasm gjson_wasm token_distribute_wasm word_count_v2_wasm word_count_wasm

clean:
	@rm -rf easyjson/easyjson.wasm

