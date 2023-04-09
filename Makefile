
_go_generate:
	docker run --rm \
		-v ${PWD}:/src \
		-w /src/go golang:1.19-alpine \
		go generate generate.go

_ts_generate:
	docker run --rm \
 		-v ${PWD}:/local \
 		openapitools/openapi-generator-cli generate \
 		-i /local/api/v1/openapi.yaml \
 		-g typescript-axios \
 		-o /local/ts/gen/v1/

openapi_generate: _go_generate _ts_generate
