GO = GO111MODULE=on GOPROXY="https://goproxy.cn,direct" go
ENTRY = main.go
RUNTIME_CONFIG_FILE = ./config/runtime.config
BINARY_NAME = ./aurora

# config
config:
	@${GO} run task/mysql.go

# compile
binary:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 ${GO} build -o ${BINARY_NAME}

# serve
development:
	${GO} run ${ENTRY} -conf ${RUNTIME_CONFIG_FILE}
server:
	${BINARY_NAME} -conf ${RUNTIME_CONFIG_FILE}

# clean
clean:
	rm ${BINARY_NAME}

# alias
package: binary
dev: development

.PHONY: config binary development server clean package dev