GOOS ?= ${shell go env GOOS}
GOARCH ?= ${shell go env GOARCH}
GOBUILD = CGO_ENABLED=0 go build
ID = ${shell id -u}

.PHONY: default
default:
	@echo ${GOOS}
	@echo ${GOARCH}
	@echo ${GOBUILD}
	@echo ${ID}

.PHONY: build
build:
	${GOBUILD} -o jaeger-bigquery-${GOOS}-${GOARCH} ./cmd/main.go

.PHONY: run
run:
	docker run --rm --name jaeger-bq JAEGER_DISABLED=false --link some-clickhouse-server -p16686:16686 -p14250:14250 -p14268:14268 -p6831:6831/udp -v "${PWD}:/data" -e SPAN_STORAGE_TYPE=grpc-plugin jaegertracing/all-in-one:${JAEGER_VERSION} --query.ui-config=/data/jaeger-ui.json --grpc-storage-plugin.binary=/data/jaeger-clickhouse-$(GOOS)-$(GOARCH) --grpc-storage-plugin.configuration-file=/data/config.yaml --grpc-storage-plugin.log-level=debug

