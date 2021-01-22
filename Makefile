GO = GO111MODULE=on GOPROXY="https://goproxy.cn,direct" go
ENTRY = main.go
RUNTIME_CONFIG_FILE = ./config/runtime.config
BINARY_PATH = ./aurora
IMAGE_NAME = qmdx00/aurora

# config
config:
	@${GO} run task/mysql.go

# compile
binary:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 ${GO} build -o ${BINARY_PATH}
	cp /usr/local/go/lib/time/zoneinfo.zip zoneinfo.zip

# develop
develop:
	${GO} run ${ENTRY} -conf ${RUNTIME_CONFIG_FILE}

# server
server: binary
	${BINARY_PATH} -conf ${RUNTIME_CONFIG_FILE}

# docker image
image: binary
	docker build -t ${IMAGE_NAME} .

# clean
clean:
	rm ${BINARY_PATH}

# help
help: binary
	${BINARY_NAME} -h

# alias
dev: develop

.PHONY: config binary develop server image clean help dev