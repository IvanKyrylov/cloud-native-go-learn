build:
	go build -o ./bin/ ./cmd/service/.

run: build
	./bin/service

lint: ## Run GoLangCI Lint
	golangci-lint run ./... --config ./.golangci.yml

#generate-swagger-api: ## generate swagger api
#	@for f in $(shell ls -d ./api/* | cut -f3 -d'/' ); do \
#  		echo "=== Generate model for $${f} ==="; \
#		rm -rf ./gen && \
#		rm -rf ./api/$${f}/openapi/ && \
#		mkdir -p ./api/$${f}/openapi/ && \
#		openapi-generator generate -g go-gin-server -i $(shell pwd)/api/$${f}/swagger.yaml -o ./gen --enable-post-process-file --additional-properties=apiPath=$${f} && \
#		mv ./gen/$${f}/model_*.go ./api/$${f}/openapi && \
#		cp ./api/$${f}/swagger.yaml ./cmd/doc-serve/static/$${f}.yaml && \
#  		echo "=== End Generate for $${f} ==="; \
#  	done
#	@rm -rf ./gen