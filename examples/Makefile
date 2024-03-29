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

kv_wasm:
	@cd kv && tinygo build -o kv.wasm -scheduler=none --no-debug -target=wasi main.go
	@mkdir -pv wasms && mv kv/kv.wasm wasms
	@echo built to wasms/kv.wasm

sql_wasm:
	@cd sql && tinygo build -o sql.wasm -scheduler=none --no-debug -target=wasi main.go
	@mkdir -pv wasms && mv sql/sql.wasm wasms
	@echo built to wasms/sql.wasm

crypto_wasm:
	@cd crypto && tinygo build -o crypto.wasm -scheduler=none --no-debug -target=wasi main.go
	@mkdir -pv wasms && mv crypto/crypto.wasm wasms
	@echo built to wasms/crypto.wasm

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

env_wasm:
	@cd env && tinygo build -o env.wasm -scheduler=none --no-debug -target=wasi main.go
	@mkdir -pv wasms && mv env/env.wasm wasms
	@echo built to wasms/env.wasm

stream_computing_wasm:
	@cd stream_computing && tinygo build -o stream_computing.wasm -scheduler=none --no-debug -target=wasi main.go
	@mkdir -pv wasms && mv stream_computing/stream_computing.wasm wasms
	@echo built to wasms/stream_computing.wasm

mqtt_wasm:
	@cd mqtt && tinygo build -o mqtt.wasm -scheduler=none --no-debug -target=wasi mqtt.go
	@mkdir -pv wasms && mv mqtt/mqtt.wasm wasms
	@echo built to wasms/mqtt.wasm

build: easyjson_wasm log_wasm gjson_wasm token_distribute_wasm kv_wasm sql_wasm word_count_wasm env_wasm stream_computing_wasm mqtt_wasm

clean:
	@rm -rf easyjson/easyjson.wasm

